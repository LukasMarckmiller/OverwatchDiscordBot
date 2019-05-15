package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//TODO Add Training Message Content as Image URL and make URL editable by !Edit !Training Command
//TODO Overwatch API integration
const (
	//Do not change
	TokenType          = "Bot"
	EventReady         = "READY"
	EventMessageCreate = "MESSAGE_CREATE"
	DiscordBaseUrl     = "https://discordapp.com/api/"
	DeviceName         = "Odroid XU4Q"
	BrowserName        = "Chromium"
	//Changeable
	DBPATH = "/home/lab01/db" //"C:\\Users\\Lukas\\go\\src\\OverwatchDiscordBot"//
)

type session struct {
	db *dbSession
	ws *websocketSession
}

var thisSession session

func main() {
	thisSession = session{}

	for {
		thisSession.ws = &websocketSession{SequenzNumber: 0}

		con, err := thisSession.ws.openCon()
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
		thisSession.db = dbs
		go startAlarmClock(6, 0, 0, pollingCustomPlayers) //Set alarm clock for polling stats to 6:00:00am (pm would be setAlarmClock(18,0,0), timezone is based on current timezone

		//func blocks
		err = thisSession.ws.startListener(con)

		if err != nil {
			fmt.Printf("Failed to listen to discord websocket connection. Fallback mechanism is trying to connect again in 5 seconds\n")
			fmt.Printf("Error:\n%v\n", err)
			_ = con.Close()
			time.Sleep(5 * time.Second)
			continue
		}
	}
}
/*
var players = [] string{
	"Exploit-21751",
	"Alex-2476",
	"trable-21221",
	"Valyria-21126",
	"Zayana-2698",
	"Ronin-23639",
	"FakeKevin-2756",
	"litchblade-2244",
	"HealMePlease-21234",
}
*/
func pollingCustomPlayers() error {
	records, err := thisSession.db.driver.ReadAll(CollectionPlayer)
	if err != nil {
		return err
	}

	var playerStats []owStatsPersistenceLayer

	for _, record := range records {
		playerStat := owStatsPersistenceLayer{}
		if err := json.Unmarshal([]byte(record), &playerStat); err != nil {
			return err
		}
		playerStats = append(playerStats, playerStat)
	}

	for _, player := range playerStats {
		var guildSettings guildSettingsPersistenceLayer
		if err = thisSession.db.getGuildConfig(player.Guild, &guildSettings); err != nil {
			return err
		}

		owPlayerStats, err := getPlayerStats(player.Battletag, guildSettings.Platform, guildSettings.Region)
		if err != nil {
			return err
		}

		var owPersLayerObj = owStatsPersistenceLayer{OWPlayer: *owPlayerStats, Battletag: player.Battletag}
		if err = thisSession.db.writePlayer(owPersLayerObj); err != nil {
			return err
		}
	}

	return nil
}
