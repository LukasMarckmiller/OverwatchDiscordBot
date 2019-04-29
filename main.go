package main

import (
	"fmt"
	"time"
)

//TODO Add Training Message Content as Image URL and make URL editable by !Edit !Training Command
//TODO Overwatch API integration
const(
	//Do not change
	TokenType          = "Bot"
	EventReady         = "READY"
	EventMessageCreate = "MESSAGE_CREATE"
	DiscordBaseUrl     = "https://discordapp.com/api/"
	DeviceName         = "Odroid XU4Q"
	BrowserName        = "Chromium"
	//Changeable
	RetryTimes = 5
)

type getCommandContent func(param string) string

func main() {

	for {
		s := Session{SequenzNumber: 0}
		con, err := s.openCon()
		var errorCnt int
		for err != nil {
			if errorCnt >= RetryTimes {
				fmt.Printf("Cant connect to discord websocket. Restarting ...\n")
			}
			errorCnt++
			fmt.Printf("Failed to open connection to discord websocket. Fallback mechanism is trying to connect again in 5 seconds\n")
			fmt.Printf("Error:\n%v", err)

			time.Sleep(5 * time.Second)
			con, err = s.openCon()
		}
		fmt.Printf("Connection to discord established. Received Hello.\n")
		err = s.startListener(con)
		errorCnt = 0
		for err != nil {
			if errorCnt >= RetryTimes {
				fmt.Printf("Cant restart Listener. Restarting ...\n")
			}
			errorCnt++
			fmt.Printf("Failed to listen to discord websocket connection. Fallback mechanism is trying to connect again in 5 seconds\n")
			fmt.Printf("Error:\n%v", err)
			time.Sleep(5 * time.Second)
			err = s.startListener(con)
		}
	}
}
