package main

import (
	"net/http"
	"fmt"
	"bufio"
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
	"runtime"
	"strings"
	"io"
	"io/ioutil"
	"bytes"
	"errors"
)


const(
	TOKEN = "NTY1MjI5NjQwNjQ2MzkzODk1.XMNDPg.jL-v1B4oikhDVwt5XpEVr_Jeg0k"
	TOKEN_TYPE = "Bot"
	EVENT_READY          = "READY"
	EVENT_MESSAGE_CREATE = "MESSAGE_CREATE"
	MESSAGE_TRAINING = "Trainings:\r\nMontag: ab 19:30 (Scrim, Review)\r\nDienstag: ab 19:30 (Scrim, Review)\r\nDonnerstag ab 19:30 (Ranked)"
	MESSAGE_HELP = "Currently supported commands:\r\n!Training : Zeigt aktuelle Trainigszeiten"
	DISCORD_BASE_URL = "https://discordapp.com/api/"

	DEVICE_NAME = "Odroid XU4Q"
	BROWSER_NAME = "Chromium"
)

type Session struct {
	 SequenzNumber int
	 HeartbeatInterval int
	 Client http.Client
}

func main() {

	s := Session{SequenzNumber:0, Client:http.Client{}}

	// WEBHOOK HANDSHAKE

	//1 GET Webhook URL
	resp, err := s.sendHTTPRequest(http.MethodGet,DISCORD_BASE_URL + "gateway/bot", nil)
	if err != nil{
		fmt.Printf("Error: %v\n", err)
	}
	reader := bufio.NewScanner(resp.Body)
	var gatewayRepresentation discordGatewayRepresentation
	for reader.Scan() {
		respString := reader.Bytes()
		if err = json.Unmarshal(respString, &gatewayRepresentation); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Printf("%+v\n", gatewayRepresentation)
	}

	//2 Connect to Websocket and Receive OpCode 10 Hello
	con, payload, err := s.websocketConnect(gatewayRepresentation.URL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Connection to discord established. Received Hello.\n")

	//3.1 Start Heartbeat
	ticker := time.NewTicker(time.Duration(s.HeartbeatInterval) * time.Millisecond)
	defer ticker.Stop()

	go func() {
		//Sent first heartbeat
		con.WriteMessage(websocket.TextMessage, []byte("{\"op\": 1,\"d\": null}"))

		for _ = range ticker.C {
			fmt.Printf("Send Heartbeat\n")

			discordHeartbeat := discordHeartbeat{Op: 1, D: s.SequenzNumber}
			data, err := json.Marshal(&discordHeartbeat)
			if err != nil {
				fmt.Printf("Error parsing heartbeat object to JSON\n")
			}

			if err := con.WriteMessage(websocket.TextMessage, data); err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}()

	//3.2 Receive First Heartbeat
	_, data, err := con.ReadMessage()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if err = json.Unmarshal(data, &payload); err != nil {
		fmt.Printf("%v\n", err)
	}

	s.SequenzNumber = payload.S

	if payload.Op != 11 {
		fmt.Printf("No first Ack received. Probably no Connection\n")
		return
	}

	//Identity
	identity := discordIdentityPayload{Op: 2, D: discordIdentity{Token: TOKEN_TYPE + " " + TOKEN, Properties: dicsordIdentityProperties{Device: DEVICE_NAME, Os: runtime.GOOS, Browser: BROWSER_NAME}}}

	data, err = json.Marshal(&identity)
	if (err != nil) {
		fmt.Printf("Error parsing identity object to JSON\n")
		return
	}
	if err = con.WriteMessage(websocket.TextMessage, data); err != nil {
		fmt.Printf("%v\n", err)
	}

	_, data, err = con.ReadMessage()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	//Catch Ready event to finish Handshake
	var event discordWebsocketPayloadPresentation
	if err = json.Unmarshal(data, &event); err != nil {
		fmt.Printf("%v\n", err)
	}
	if (event.T != EVENT_READY) {
		fmt.Printf(fmt.Sprintf("Failed to idenitfy at discord websocket server. Expected READY but got %v\n", event.T))
	}

	//Basic Listener
	for {
		//TODO check acks , start timer when sending heartbeat, stop timer when con.ReadMessage Op=11, when timer expires increment missedAck counter
		//TODO when miseedAckCounter has certain value then restart
		_, data, err = con.ReadMessage()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		var event discordWebsocketPayloadPresentation
		if err = json.Unmarshal(data, &event); err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%+v\n", event.T)
		s.SequenzNumber = event.S

		//Handle event
		switch event.T {
		case EVENT_MESSAGE_CREATE:
			var messagePayload discordMessageObject
			if err = json.Unmarshal(event.D, &messagePayload); err != nil {
				fmt.Printf("%v\n", err)
			}

			switch strings.Split(messagePayload.Content, " ")[0] {
			case "!Training":
				resp, err := s.sendMessageToChannel(MESSAGE_TRAINING,messagePayload.ChannelId)
				if err != nil{
					fmt.Printf("%v\n", err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
				fmt.Printf("%v\nBody:\n%s", resp.Status, body)
			case "!Help":
				resp, err := s.sendMessageToChannel(MESSAGE_HELP,messagePayload.ChannelId)
				if err != nil{
					fmt.Printf("%v\n", err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
				fmt.Printf("%v\nBody:\n%s", resp.Status, body)
			default:
				break
			}
		}
	}
}

func (s *Session)sendMessageToChannel(content string, channelId string) (*http.Response, error) {
	responseMessage := discordMessageRequest{Content: content, Tts: false}
	data, err := json.Marshal(responseMessage)
	if err != nil {
		return nil, err
	}
	resp, err := s.sendHTTPRequest(http.MethodPost, fmt.Sprintf("%v/channels/%v/messages", DISCORD_BASE_URL, channelId), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (s *Session)sendHTTPRequest(method string,URL string,body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", TOKEN_TYPE+ " " + TOKEN)
	req.Header.Add("Content-type", "application/json")
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprint("HTTP Ok (200) expected, but got %v",resp.StatusCode))
	}

	return resp, err
}

func (s *Session)websocketConnect (websocketURL string, ) (*websocket.Conn,discordWebsocketPayloadPresentation, error) {
	payload := discordWebsocketPayloadPresentation{}
	helloPayload := discordWebsocketHelloPresentation{}
	con, _, err := websocket.DefaultDialer.Dial(websocketURL, nil)
	if err != nil {
		return nil,payload, err
	}

	_, data, err := con.ReadMessage()
	if err != nil {
		return nil,payload, err
	}

	if err = json.Unmarshal(data, &payload); err != nil {
		return nil,payload, err
	}

	if payload.Op != 10 {
		return nil,payload, errors.New("Op code was not 10 after websocket init\n")
	}


	if err = json.Unmarshal([]byte(payload.D), &helloPayload); err != nil {
		return nil,payload, err
	}
	s.HeartbeatInterval = helloPayload.HeartbeatInterval

	return con,payload, err
}
