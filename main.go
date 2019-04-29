package main

import (
	"fmt"
	"time"
)

//TODO Add Training Message Content as Image URL and make URL editable by !Edit !Training Command
//TODO Overwatch API integration
const(
	TOKEN_TYPE           = "Bot"
	EVENT_READY          = "READY"
	EVENT_MESSAGE_CREATE = "MESSAGE_CREATE"
	DISCORD_BASE_URL     = "https://discordapp.com/api/"
	DEVICE_NAME  = "Odroid XU4Q"
	BROWSER_NAME = "Chromium"
	RETRY_TIMES  = 5

	COMMAND_HELP     = "Currently supported commands:\r\n!Training : Zeigt aktuelle Trainigszeiten"
	COMMAND_TRAINING = "Trainings:\r\nMontag: ab 19:30 (Scrim, Review)\r\nDienstag: ab 19:30 (Scrim, Review)\r\nDonnerstag ab 19:30 (Ranked)"
)

type getCommandContent func(param string) string

func main() {

	for {
		s := Session{SequenzNumber: 0}
		con, err := s.openCon()
		var errorCnt int
		for err != nil {
			if errorCnt >= RETRY_TIMES {
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
			if errorCnt >= RETRY_TIMES {
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
