package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

const (
	//TokenType specifies the authentication user
	TokenType = "Bot"
	//EventReady - Event Ready
	EventReady = "READY"
	//EventMessageCreate - is send when a message is posted to a channel
	EventMessageCreate = "MESSAGE_CREATE"
	//EventGuildCreate - is send when connection to a guild is established
	EventGuildCreate = "GUILD_CREATE"
	//DiscordBaseUrl - Base url for discord rest api
	DiscordBaseUrl = "https://discordapp.com/api"
	//DiscordBaseImageUrl - Base url for discord statics
	DiscordBaseImageUrl = "https://cdn.discordapp.com"
	//DeviceName - Used for activity feed
	DeviceName = "Odroid XU4Q"
	//BrowserName - Used for activity feed
	BrowserName = "Chromium"
)

var (
	client http.Client
)

type websocketSession struct {
	SequenzNumber     int
	HeartbeatInterval int
	BotUserId         string
	events            events
}

func (s *websocketSession) openCon() (con *websocket.Conn, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("[panic in startListener]: %s\n[!stack]: %s \n", r, debug.Stack()))
			return
		}
	}()

	// WEBHOOK HANDSHAKE

	//1 GET Webhook URL
	resp, err := s.sendHTTPDiscordRequest(http.MethodGet, DiscordBaseUrl+"/gateway/bot", nil)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewScanner(resp.Body)
	var gatewayRepresentation discordGatewayRepresentation
	for reader.Scan() {
		respString := reader.Bytes()
		if err = json.Unmarshal(respString, &gatewayRepresentation); err != nil {
			return nil, err
		}

		fmt.Printf("Connected to discord web api and got websocket URL.\n")
	}

	//2 Connect to Websocket and Receive OpCode 10 Hello
	con, payload, err := s.websocketConnect(gatewayRepresentation.URL)
	if err != nil {
		return nil, err
	}

	//Sent first heartbeat
	if err = con.WriteMessage(websocket.TextMessage, []byte("{\"op\": 1,\"d\": null}")); err != nil {
		return nil, err
	}

	//3.2 Receive First Heartbeat
	_, data, err := con.ReadMessage()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}

	s.SequenzNumber = payload.S

	if payload.Op != 11 {
		return nil, fmt.Errorf("Expected Opcode 11 as answer to first heartbeat but got %v\n", payload.Op)
	}

	//Identity
	identity := discordIdentityPayload{Op: 2, D: discordIdentity{Token: TokenType + " " + TOKEN, Properties: dicsordIdentityProperties{Device: DeviceName, Os: runtime.GOOS, Browser: BrowserName}}}

	data, err = json.Marshal(&identity)
	if err != nil {
		return nil, err
	}
	if err = con.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, err
	}

	_, data, err = con.ReadMessage()
	if err != nil {
		return nil, err
	}
	//Catch Ready event to finish Handshake
	var event discordWebsocketPayloadPresentation
	if err = json.Unmarshal(data, &event); err != nil {
		return nil, err
	}
	if event.T != EventReady {
		return nil, fmt.Errorf("Failed to identify at discord websocket server. Expected READY but got %v\n", event.T)
	}
	var eventPayload discordReadyEventObject
	if err = json.Unmarshal(event.D, &eventPayload); err != nil {
		return nil, err
	}
	s.BotUserId = eventPayload.User.Id

	//Set bot activity
	var presenceUpdate = discordWebsocketPayloadPresentation{Op: 3,
		D: []byte(fmt.Sprintf("{ \"since\": %v, \"game\": { \"name\": \"Golang on %v\", \"type\": 0 }, \"status\": \"online\", \"afk\": false }",
			time.Now().UnixNano()/int64(time.Millisecond),
			DeviceName))}

	data, err = json.Marshal(&presenceUpdate)
	if err != nil {
		return nil, err
	}

	if err = con.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, err
	}

	return con, nil
}

func (s *websocketSession) startListener(con *websocket.Conn) (error error) {
	//Recover on panic
	defer func() {
		if r := recover(); r != nil {
			error = errors.New(fmt.Sprintf("[panic in startListener]: %s\n[!stack]: %s \n", r, debug.Stack()))
			return
		}
	}()

	//3.1 Start Heartbeat
	ticker := time.NewTicker(time.Duration(s.HeartbeatInterval) * time.Millisecond)
	defer ticker.Stop()
	defer con.Close()

	go func() {
		for range ticker.C {
			discordHeartbeat := discordHeartbeat{Op: 1, D: s.SequenzNumber}
			data, err := json.Marshal(&discordHeartbeat)
			if err != nil {
				fmt.Printf("%v\n", err)
			}

			if err := con.WriteMessage(websocket.TextMessage, data); err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}()
	fmt.Printf("Listener started.\n")
	//Basic Listener
	for {
		//TODO check acks , start timer when sending heartbeat, stop timer when con.ReadMessage Op=11, when timer expires increment missedAck counter
		//TODO when miseedAckCounter has certain value then restart
		_, data, err := con.ReadMessage()
		if err != nil {
			return err
		}

		var event discordWebsocketPayloadPresentation
		if err = json.Unmarshal(data, &event); err != nil {
			return err
		}
		s.SequenzNumber = event.S

		//Handle event
		s.events = events{}
		switch event.T {
		case EventMessageCreate:
			if err = s.events.handleMessageCreate(event); err != nil {
				return err
			}

		case EventGuildCreate:
			var guild discordGuildObject
			if err = json.Unmarshal(event.D, &guild); err != nil {
				return err
			}
			fmt.Printf("%+v Opcode: %v Guild: %s | %s  \n", event.T, event.Op, guild.Name, guild.Id)
		}
	}
}

func (s *websocketSession) sendMessageToChannel(content discordMessageRequest, channelId string) (discordMessageResponse, error) {
	//responseMessage := discordMessageRequest{Content: content, Tts: false}
	//As long as not specifically important, disable global
	var responseMessage discordMessageResponse
	content.Tts = false
	data, err := json.Marshal(content)
	if err != nil {
		return responseMessage, err
	}
	resp, err := s.sendHTTPDiscordRequest(http.MethodPost, fmt.Sprintf("%v/channels/%v/messages", DiscordBaseUrl, channelId), bytes.NewReader(data))
	if err != nil {
		return responseMessage, err
	}

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.Body)
	_ = json.Unmarshal(buf.Bytes(), &responseMessage)
	return responseMessage, err
}

func (s *websocketSession) updateMessageInChanel(content discordMessageRequest, channelId string, messageId string) error {
	data, err := json.Marshal(content)
	if err != nil {
		return err
	}
	_, err = s.sendHTTPDiscordRequest(http.MethodPatch, fmt.Sprintf("%v/channels/%v/messages/%v", DiscordBaseUrl, channelId, messageId), bytes.NewReader(data))
	if err != nil {
		return err
	}

	return err
}

func (s *websocketSession) triggerTypingInChannel(channelId string) (*http.Response, error) {
	resp, err := s.sendHTTPDiscordRequest(http.MethodPost, fmt.Sprintf("%v/channels/%v/typing", DiscordBaseUrl, channelId), nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *websocketSession) getUserAvatarOrDefaultUrl(userId string, avatarHash string, userDiscriminator string) (url string, err error) {
	//Try to get user avatar
	url = fmt.Sprintf("%s/avatars/%s/%s.webp", DiscordBaseImageUrl, userId, avatarHash)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "image/webp")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusOK {
		return

	} else if resp.StatusCode == http.StatusNotFound {
		usrDiscInt, _ := strconv.Atoi(userDiscriminator)
		url = fmt.Sprintf("%s/embed/avatars/%d.png", DiscordBaseImageUrl, usrDiscInt%5)
		return
	} else {
		return "", errors.New("Got Status while trying to get the users avatar." + resp.Status)
	}
}

func (s *websocketSession) sendHTTPDiscordRequest(method string, URL string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", TokenType+" "+TOKEN)
	req.Header.Add("Content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != 204 {
		return nil, fmt.Errorf("HTTP Ok (200) expected, but got %v url: %v %v", resp.StatusCode, method, URL)
	}

	return resp, err
}

func (s *websocketSession) websocketConnect(websocketURL string) (*websocket.Conn, discordWebsocketPayloadPresentation, error) {
	payload := discordWebsocketPayloadPresentation{}
	helloPayload := discordWebsocketHelloPresentation{}
	con, _, err := websocket.DefaultDialer.Dial(websocketURL, nil)
	if err != nil {
		return nil, payload, err
	}

	_, data, err := con.ReadMessage()
	if err != nil {
		return nil, payload, err
	}

	if err = json.Unmarshal(data, &payload); err != nil {
		return nil, payload, err
	}

	if payload.Op != 10 {
		return nil, payload, errors.New("Op code was not 10 after websocket init\n")
	}

	if err = json.Unmarshal([]byte(payload.D), &helloPayload); err != nil {
		return nil, payload, err
	}
	s.HeartbeatInterval = helloPayload.HeartbeatInterval

	return con, payload, err
}
