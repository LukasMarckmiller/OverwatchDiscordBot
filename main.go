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
	DBPATH = "/home/lab01/db" //"C:\\Users\\Lukas\\go\\src\\OverwatchDiscordBot\\db"
)

type getCommandContent func(param string) string

var db *dbSession

func main() {
	for {
		s := Session{SequenzNumber: 0}

		con, err := s.openCon()
		if err != nil {
			fmt.Printf("Failed to open connection to discord websocket. Fallback mechanism is trying to connect again in 5 seconds\n")
			fmt.Printf("Error:\n%v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		fmt.Printf("Connection to discord established. Received Hello.\n")

		dbs, err := createDB(DBPATH)
		if err != nil {
			fmt.Println(err)
			break
		}
		db = dbs

		go startAlarmClock(6, 0, 0, pollingCustomPlayers) //Set alarm clock for polling stats to 6:00:00am (pm would be setAlarmClock(18,0,0), timezone is based on current timezone

		//func blocks
		err = s.startListener(con)

		if err != nil {
			fmt.Printf("Failed to listen to discord websocket connection. Fallback mechanism is trying to connect again in 5 seconds\n")
			fmt.Printf("Error:\n%v\n", err)
			con.Close()
			time.Sleep(5 * time.Second)
			continue
		}
	}
}

var players = [] string{
	"Exploit-21751",
	"Alex-2476",
	"trable-21221",
	"Valyria-21126",
	"Zayana-2698",
	"Ronin-23639",
	"FakeKevin-2756",
}

func pollingCustomPlayers() error {
	for _, player := range players {
		owPlayerStats, err := getPlayerStats(player)
		if err != nil {
			return err
		}
		var owPersLayerObj = owStatsPersistenceLayer{OWPlayer: *owPlayerStats, Battletag: player}
		if err = db.write(owPersLayerObj); err != nil {
			return err
		}
	}

	return nil
}

