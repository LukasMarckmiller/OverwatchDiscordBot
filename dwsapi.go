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
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	Client http.Client
)

type websocketSession struct {
	SequenzNumber        int
	HeartbeatInterval    int
	BotUserId            string
	cachedMessagePayload discordMessageObject
}

func (s *websocketSession) openCon() (*websocket.Conn, error) {

	// WEBHOOK HANDSHAKE

	//1 GET Webhook URL
	resp, err := s.sendHTTPDiscordRequest(http.MethodGet, DiscordBaseUrl+"gateway/bot", nil)
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
		return nil, errors.New(fmt.Sprintf("Expected Opcode 11 as answer to first heartbeat but got %v\n", payload.Op))
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
		return nil, errors.New(fmt.Sprintf("Failed to idenitfy at discord websocket server. Expected READY but got %v\n", event.T))
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

func (s *websocketSession) startListener(con *websocket.Conn) error {

	//3.1 Start Heartbeat
	ticker := time.NewTicker(time.Duration(s.HeartbeatInterval) * time.Millisecond)
	defer ticker.Stop()
	defer con.Close()

	go func() {
		//TODO Add chanel to communicate with the heartbeat routine and send value to channel if error
		//Check channel in non blocking select block in Basic listener and abort if value in channel

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
		fmt.Printf("%+v Opcode: %v\n", event.T, event.Op)
		s.SequenzNumber = event.S

		//Handle event
		switch event.T {
		case EventMessageCreate:

			if err = json.Unmarshal(event.D, &s.cachedMessagePayload); err != nil {
				return err
			}
			//Filter non command
			if !strings.HasPrefix(s.cachedMessagePayload.Content, "!") {
				break
			}
			//Filter if requesting event triggered by this bot
			if s.cachedMessagePayload.Author.Id == s.BotUserId {
				break
			}
			_, err := s.triggerTypingInChannel(s.cachedMessagePayload.ChannelId)
			if err != nil {
				return err
			}

			command := strings.Split(s.cachedMessagePayload.Content, " ")[0]
			var message string
			cmd, ok := commandMap[command]
			if !ok {
				message = "Command not supported: " + command
				break
			}

			content := strings.Trim(strings.Replace(s.cachedMessagePayload.Content, command, "", -1), " ")
			regex, err := regexp.Compile("\".*\"")
			// i.e !Command "here is a mulit param" ,  !Command ThisIsASingleParam
			//Find multiword comment
			multiParam := regex.FindString(content)
			//Remove it
			content = strings.Replace(content, multiParam, "", 1)

			if multiParam, _ = strconv.Unquote(multiParam); err != nil {
				//No multi word param was send by the client
			}

			if content != "" {
				params := strings.Split(content, " ")
				if multiParam != "" {
					params = append(params, multiParam)
				}
				message = cmd(params)

			} else {
				message = cmd(nil)

			}
			_, err = s.sendMessageToChannel(message, s.cachedMessagePayload.ChannelId)
			if err != nil {
				return err
			}
		}
	}
}

func (s *websocketSession) sendMessageToChannel(content string, channelId string) (*http.Response, error) {
	responseMessage := discordMessageRequest{Content: content, Tts: false}
	data, err := json.Marshal(responseMessage)
	if err != nil {
		return nil, err
	}
	resp, err := s.sendHTTPDiscordRequest(http.MethodPost, fmt.Sprintf("%v/channels/%v/messages", DiscordBaseUrl, channelId), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (s *websocketSession) triggerTypingInChannel(channelId string) (*http.Response, error) {
	resp, err := s.sendHTTPDiscordRequest(http.MethodPost, fmt.Sprintf("%v/channels/%v/typing", DiscordBaseUrl, channelId), nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *websocketSession) sendHTTPDiscordRequest(method string, URL string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", TokenType+" "+TOKEN)
	req.Header.Add("Content-type", "application/json")
	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != 204 {
		return nil, errors.New(fmt.Sprintf("HTTP Ok (200) expected, but got %v", resp.StatusCode))
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
