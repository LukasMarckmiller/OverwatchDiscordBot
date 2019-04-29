package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var (
	commandMap = map[string]getCommandContent{
		"!Training": getTrainingTimes,
		"!Help":     getCurrentlySupportedCommands,
		"!Stats":    getOverwatchPlayerStats,
	}
)

func getTrainingTimes(param string) string {
	//param unused
	return COMMAND_TRAINING
}

func getCurrentlySupportedCommands(param string) string {
	//param unused
	return COMMAND_HELP
}

func getOverwatchPlayerStats(param string) string {
	//TODO Check if username seperated by - instead of #
	requ, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://ow-api.com/v1/stats/pc/eu/%v/profile", param), nil)
	if err != nil {
		return "An error while retrieving data from the Overwatch stats api occured.\n" + err.Error()
	}
	resp, err := Client.Do(requ)
	if err != nil {
		return "An error while retrieving data from the Overwatch stats api occured.\n" + err.Error()
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		return "An error while reading the response from the Overwatch API, player request.\n" + err.Error()
	}
	var owPlayerStats OWPlayer
	err = json.Unmarshal(bytes, &owPlayerStats)
	return fmt.Sprintf("Statistik für Spieler: %v\nRating: %v\nCompetitive Games played: %v Games won: %v\n",
		owPlayerStats.Name,
		owPlayerStats.Rating,
		owPlayerStats.CompetitiveStats.Games.Played,
		owPlayerStats.CompetitiveStats.Games.Won)
}
