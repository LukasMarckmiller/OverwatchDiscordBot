package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//TODO Add Training Message Content as Image URL and make URL editable by !Edit !Training Command
//TODO Overwatch API integration
const (
	//Changeable
	DBPATH  = "/home/lab01/db" //"C:\\Users\\Lukas\\go\\src\\OverwatchDiscordBot"//
	VERSION = "1.0.3"
)

type session struct {
	db *dbSession
	ws *websocketSession
}

var thisSession session

func main() {
	fmt.Println("###### OverwatchTeamBot Current Version: " + VERSION + "######")
	for {
		thisSession = session{}
		onError := make(chan bool)

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
		startAlarmClock(8, 0, 0, pollingCustomPlayers, onError) //Set alarm clock for polling stats to 6:00:00am (pm would be setAlarmClock(18,0,0), timezone is based on current timezone

		err = thisSession.ws.startListener(con)

		if err != nil {
			fmt.Printf("Failed to listen to discord websocket connection. Fallback mechanism is trying to connect again in 5 seconds\n")
			fmt.Printf("Error:\n%v\n", err)
			//Close con
			_ = con.Close()
			//Close alarm go routine
			close(onError)
			time.Sleep(5 * time.Second)
			continue
		}
	}
}

func pollingCustomPlayers() error {
	records, err := thisSession.db.driver.ReadAll(collectionPlayer)
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
		platform := player.Platform
		region := player.Region

		owPlayerStats, err := getPlayerStats(player.Battletag, platform, region)
		if err != nil {
			return err
		}

		var owPersLayerObj = owStatsPersistenceLayer{OWPlayer: *owPlayerStats, Battletag: player.Battletag, Guild: player.Guild, Platform: platform, Region: region}
		if err = thisSession.db.writePlayer(owPersLayerObj); err != nil {
			return err
		}
	}

	return nil
}
