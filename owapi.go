package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type OWPlayer struct {
	CompetitiveStats struct {
		Awards struct {
			Cards        int `json:"cards"`
			Medals       int `json:"medals"`
			MedalsBronze int `json:"medalsBronze"`
			MedalsSilver int `json:"medalsSilver"`
			MedalsGold   int `json:"medalsGold"`
		} `json:"awards"`
		Games struct {
			Played int `json:"played"`
			Won    int `json:"won"`
		} `json:"games"`
	} `json:"competitiveStats"`
	Endorsement     int    `json:"endorsement"`
	EndorsementIcon string `json:"endorsementIcon"`
	GamesWon        int    `json:"gamesWon"`
	Icon            string `json:"icon"`
	Level           int    `json:"level"`
	LevelIcon       string `json:"levelIcon"`
	Name            string `json:"name"`
	Prestige        int    `json:"prestige"`
	PrestigeIcon    string `json:"prestigeIcon"`
	Private         bool   `json:"private"`
	QuickPlayStats struct {
		Awards struct {
			Cards        int `json:"cards"`
			Medals       int `json:"medals"`
			MedalsBronze int `json:"medalsBronze"`
			MedalsSilver int `json:"medalsSilver"`
			MedalsGold   int `json:"medalsGold"`
		} `json:"awards"`
		Games struct {
			Played int `json:"played"`
			Won    int `json:"won"`
		} `json:"games"`
	} `json:"quickPlayStats"`
	Rating     int    `json:"rating"`
	RatingIcon string `json:"ratingIcon"`
}

func getPlayerStats(battletag string) (*OWPlayer, error) {
	requ, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://ow-api.com/v1/stats/pc/eu/%v/profile", battletag), nil)
	if err != nil {
		return nil, errors.New("An error while retrieving data from the Overwatch stats api occured.\n" + err.Error())
	}
	resp, err := Client.Do(requ)
	if err != nil {
		return nil, errors.New("An error while retrieving data from the Overwatch stats api occured.\n" + err.Error())
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("An error while reading the response from the Overwatch API, player request.\n" + err.Error())
	}
	var owPlayerStats OWPlayer
	err = json.Unmarshal(bytes, &owPlayerStats)
	if err != nil {
		return nil, err
	}
	return &owPlayerStats, nil
}
