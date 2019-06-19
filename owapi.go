package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*type OWPlayer struct {
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
	QuickPlayStats  struct {
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
} */
type heroCareerStats struct {
	Assists struct {
		DefensiveAssists            int     `json:"defensiveAssists"`
		DefensiveAssistsAvgPer10Min float64 `json:"defensiveAssistsAvgPer10Min"`
		DefensiveAssistsMostInGame  int     `json:"defensiveAssistsMostInGame"`
		HealingDone                 int     `json:"healingDone"`
		HealingDoneAvgPer10Min      float64 `json:"healingDoneAvgPer10Min"`
		HealingDoneMostInGame       int     `json:"healingDoneMostInGame"`
		OffensiveAssists            int     `json:"offensiveAssists"`
		OffensiveAssistsAvgPer10Min float64 `json:"offensiveAssistsAvgPer10Min"`
		OffensiveAssistsMostInGame  int     `json:"offensiveAssistsMostInGame"`
	} `json:"assists"`
	Average struct {
		AllDamageDoneAvgPer10Min     float64 `json:"allDamageDoneAvgPer10Min"`
		BarrierDamageDoneAvgPer10Min float64 `json:"barrierDamageDoneAvgPer10Min"`
		DeathsAvgPer10Min            float64 `json:"deathsAvgPer10Min"`
		EliminationsAvgPer10Min      float64 `json:"eliminationsAvgPer10Min"`
		EliminationsPerLife          float64 `json:"eliminationsPerLife"`
		FinalBlowsAvgPer10Min        float64 `json:"finalBlowsAvgPer10Min"`
		HeroDamageDoneAvgPer10Min    float64 `json:"heroDamageDoneAvgPer10Min"`
		ObjectiveKillsAvgPer10Min    float64 `json:"objectiveKillsAvgPer10Min"`
		ObjectiveTimeAvgPer10Min     string  `json:"objectiveTimeAvgPer10Min"`
		SoloKillsAvgPer10Min         float64 `json:"soloKillsAvgPer10Min"`
		TimeSpentOnFireAvgPer10Min   string  `json:"timeSpentOnFireAvgPer10Min"`
	} `json:"average"`
	Best struct {
		AllDamageDoneMostInGame     int    `json:"allDamageDoneMostInGame"`
		AllDamageDoneMostInLife     int    `json:"allDamageDoneMostInLife"`
		BarrierDamageDoneMostInGame int    `json:"barrierDamageDoneMostInGame"`
		EliminationsMostInGame      int    `json:"eliminationsMostInGame"`
		EliminationsMostInLife      int    `json:"eliminationsMostInLife"`
		FinalBlowsMostInGame        int    `json:"finalBlowsMostInGame"`
		HeroDamageDoneMostInGame    int    `json:"heroDamageDoneMostInGame"`
		HeroDamageDoneMostInLife    int    `json:"heroDamageDoneMostInLife"`
		KillsStreakBest             int    `json:"killsStreakBest"`
		ObjectiveKillsMostInGame    int    `json:"objectiveKillsMostInGame"`
		ObjectiveTimeMostInGame     string `json:"objectiveTimeMostInGame"`
		SoloKillsMostInGame         int    `json:"soloKillsMostInGame"`
		TimeSpentOnFireMostInGame   string `json:"timeSpentOnFireMostInGame"`
		WeaponAccuracyBestInGame    string `json:"weaponAccuracyBestInGame"`
	} `json:"best"`
	Combat struct {
		BarrierDamageDone  int    `json:"barrierDamageDone"`
		DamageDone         int    `json:"damageDone"`
		Deaths             int    `json:"deaths"`
		Eliminations       int    `json:"eliminations"`
		FinalBlows         int    `json:"finalBlows"`
		HeroDamageDone     int    `json:"heroDamageDone"`
		ObjectiveKills     int    `json:"objectiveKills"`
		ObjectiveTime      string `json:"objectiveTime"`
		QuickMeleeAccuracy string `json:"quickMeleeAccuracy"`
		SoloKills          int    `json:"soloKills"`
		TimeSpentOnFire    string `json:"timeSpentOnFire"`
		WeaponAccuracy     string `json:"weaponAccuracy"`
	} `json:"combat"`
	Deaths       interface{}     `json:"deaths"`
	HeroSpecific json.RawMessage `json:"heroSpecific"`
	Game         struct {
		GamesLost     int    `json:"gamesLost"`
		GamesPlayed   int    `json:"gamesPlayed"`
		GamesWon      int    `json:"gamesWon"`
		TimePlayed    string `json:"timePlayed"`
		WinPercentage string `json:"winPercentage"`
	} `json:"game"`
	MatchAwards struct {
		Cards        int `json:"cards"`
		Medals       int `json:"medals"`
		MedalsBronze int `json:"medalsBronze"`
		MedalsGold   int `json:"medalsGold"`
		MedalsSilver int `json:"medalsSilver"`
	} `json:"matchAwards"`
	Miscellaneous struct {
		TurretsDestroyed int `json:"turretsDestroyed"`
	} `json:"miscellaneous"`
}

type topHeroStats struct {
	TimePlayed          string  `json:"timePlayed"`
	TimePlayedInSeconds int     `json:"timePlayedInSeconds"`
	GamesWon            int     `json:"gamesWon"`
	WinPercentage       int     `json:"winPercentage"`
	WeaponAccuracy      int     `json:"weaponAccuracy"`
	EliminationsPerLife float64 `json:"eliminationsPerLife"`
	MultiKillBest       int     `json:"multiKillBest"`
	ObjectiveKills      int     `json:"objectiveKills"`
}

type owCompleteStats struct {
	CompetitiveStats struct {
		Awards struct {
			Cards        int `json:"cards"`
			Medals       int `json:"medals"`
			MedalsBronze int `json:"medalsBronze"`
			MedalsSilver int `json:"medalsSilver"`
			MedalsGold   int `json:"medalsGold"`
		} `json:"awards"`
		CareerStats json.RawMessage `json:"careerStats"`
		Games       struct {
			Played int `json:"played"`
			Won    int `json:"won"`
		} `json:"games"`
		TopHeroes json.RawMessage `json:"topHeroes"`
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
	Rating          int    `json:"rating"`
	RatingIcon      string `json:"ratingIcon"`
}

func getPlayerStats(battletag string, platform string, region string) (*owCompleteStats, error) {

	//Prepare endpoint url part
	if region != "" {
		region = "/" + region
	}
	url := fmt.Sprintf("https://ow-api.com/v1/stats/%v%v/%v/complete", platform, region, battletag)
	requ, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.New("An error while retrieving data from the Overwatch stats api occurred\n" + err.Error())
	}

	resp, err := client.Do(requ)
	if err != nil {
		return nil, errors.New("An error while retrieving data from the Overwatch stats api occurred\n" + err.Error())
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("An error while reading the response from the Overwatch API, player request.\n" + err.Error())
	}

	var owPlayerStats owCompleteStats
	err = json.Unmarshal(bytes, &owPlayerStats)
	if err != nil {
		return nil, err
	}

	return &owPlayerStats, nil
}
