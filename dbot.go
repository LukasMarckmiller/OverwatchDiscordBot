package main

import (
	"fmt"
	"strings"
)

const (
	CommandHelp = "**Currently supported commands:**\n\t\t\t\u0060!Training\u0060: Zeigt aktuelle Trainigszeiten" +
		"\n\t\t\t\u0060!Stats <battletag>\u0060: Spieler Statistiken (z.B. *!Stats Krusher-9911*)" +
		"\n\t\t\t\u0060!Register <battletag>\u0060: Registriert neuen Spieler (z.B. *!Register Krusher-9911*)" +
		"\n\t\t\t\u0060!Update <battletag>\u0060: Aktualisiert Statistik für angegebenen Spieler (z.B. *!Update Krusher-9911*)"
	CommandTraining = "**Trainings**:\r\nMontag: ab 19:30 (Scrim, Review)\r\nDienstag: ab 19:30 (Scrim, Review)\r\nDonnerstag ab 19:30 (Ranked)"
)

var (
	commandMap = map[string]getCommandContent{
		"!Training": getTrainingTimes,
		"!Help":     getCurrentlySupportedCommands,
		"!Stats":    getOverwatchPlayerStats,
		"!Register": setNewOverwatchPlayer,
		"!Update":   setNewOverwatchPlayer,
	}
)

type getCommandContent func(param string) string

//noinspection GoUnusedParameter
func getTrainingTimes(param string) string {
	//param unused
	return CommandTraining
}

//noinspection GoUnusedParameter
func getCurrentlySupportedCommands(param string) string {
	//param unused
	return CommandHelp
}

func getOverwatchPlayerStats(param string) string {
	param = strings.Replace(param, "#", "-", 1)
	owPlayerLiveStats, err := getPlayerStats(param)
	if err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}
	var owPlayerPersistenceStats OWPlayer
	if err = s.db.readPlayer(param, &owPlayerPersistenceStats); err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}

	return fmt.Sprintf("Statistik für Spieler: **%v**\nRating: **%v**\nCompetitive Games played (all): *%v* Games won (all): *%v*\nTrend: *%d*sr (started today at *%v*)\nGames played today: *%v*\nGames won today: *%v*\n",
		owPlayerLiveStats.Name,
		owPlayerLiveStats.Rating,
		owPlayerLiveStats.CompetitiveStats.Games.Played,
		owPlayerLiveStats.CompetitiveStats.Games.Won,
		owPlayerLiveStats.Rating-owPlayerPersistenceStats.Rating,
		owPlayerPersistenceStats.Rating,
		owPlayerLiveStats.CompetitiveStats.Games.Played-owPlayerPersistenceStats.CompetitiveStats.Games.Played,
		owPlayerLiveStats.CompetitiveStats.Games.Won-owPlayerPersistenceStats.CompetitiveStats.Games.Won,
	)
}

func setNewOverwatchPlayer(param string) string {
	param = strings.Replace(param, "#", "-", 1)
	owPlayerLiveStats, err := getPlayerStats(param)
	if err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}
	owStatsPersistenceLayer := owStatsPersistenceLayer{Battletag: param, OWPlayer: *owPlayerLiveStats}
	if err = s.db.writePlayer(owStatsPersistenceLayer); err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}
	return fmt.Sprintf("Player **%v** added/refreshed.", param)
}

