package main

import (
	"fmt"
	"strings"
)

const (
	DiscordMarkupHelpURL = "https://gist.github.com/Almeeida/41a664d8d5f3a8855591c2f1e0e07b19"
	CommandHelp          = "**Currently supported commands: :information_source:**\n\t\t\t\u0060!Training\u0060: Zeigt aktuelle Trainigszeiten" +
		"\n\t\t\t\u0060!Training <value>\u0060: Aktualisiert Trainingszeiten (z.B. *!Training \"our **new** trainings are ...\"*), Fett oder Kursive Schrift? Check out Discord Markup\n\t\t\t:arrow_right:" + DiscordMarkupHelpURL +
		"\n\t\t\t\u0060!Stats <battletag>\u0060: Spieler Statistiken (z.B. *!Stats Krusher-9911*)" +
		"\n\t\t\t\u0060!Register <battletag>\u0060: Registriert neuen Spieler (z.B. *!Register Krusher-9911*)" +
		"\n\t\t\t\u0060!Update <battletag>\u0060: Aktualisiert Statistik für angegebenen Spieler (z.B. *!Update Krusher-9911*)"
	//CommandTraining = "**Trainings**:\r\nMontag: ab 19:30 (Scrim, Review)\r\nDienstag: ab 19:30 (Scrim, Review)\r\nDonnerstag ab 19:30 (Ranked)"
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

type getCommandContent func(params []string) string

//noinspection GoUnusedParameter
func getTrainingTimes(params []string) string {
	//Save param as new Training Content in DB
	if params != nil {
		if err := thisSession.db.updateTrainingDates(thisSession.ws.cachedMessagePayload.GuildId, trainingDatesPersistenceLayer{params[1]}); err != nil {
			return fmt.Sprintf("Error updating Training dates: **%v**\n*%v*\n", params[1], string(err.Error()))
		}
		return fmt.Sprintf("Training dates set to:\n*%v*", params[1])
	}
	var dates trainingDatesPersistenceLayer
	if err := thisSession.db.getTrainingDates(thisSession.ws.cachedMessagePayload.GuildId, &dates); err != nil {
		return fmt.Sprintf("Error while retrieving training dates:\n*%v*\n", string(err.Error()))
	}

	return dates.Value
}

//noinspection GoUnusedParameter
func getCurrentlySupportedCommands(params []string) string {
	//param unused
	return CommandHelp
}

func getOverwatchPlayerStats(params []string) string {
	param := strings.Replace(params[1], "#", "-", 1)
	owPlayerLiveStats, err := getPlayerStats(param)
	if err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}
	var owPlayerPersistenceStats OWPlayer
	if err = thisSession.db.readPlayer(param, &owPlayerPersistenceStats); err != nil {
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

func setNewOverwatchPlayer(params []string) string {
	param := strings.Replace(params[1], "#", "-", 1)
	owPlayerLiveStats, err := getPlayerStats(param)
	if err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}
	owStatsPersistenceLayer := owStatsPersistenceLayer{Battletag: param, OWPlayer: *owPlayerLiveStats}
	if err = thisSession.db.writePlayer(owStatsPersistenceLayer); err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}
	return fmt.Sprintf("Player **%v** added/refreshed.", param)
}

