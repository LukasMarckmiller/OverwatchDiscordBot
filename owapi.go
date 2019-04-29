package main

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
