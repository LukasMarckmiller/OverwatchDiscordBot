package main

import (
	"fmt"
	"github.com/revel/cmd/utils"
	"strings"
)

const (
	DiscordMarkupHelpURL = "https://gist.github.com/Almeeida/41a664d8d5f3a8855591c2f1e0e07b19"
	CommandHelp          = "**Currently supported commands: :information_source:**\n\t\t\t\u0060!Training\u0060: Zeigt aktuelle Trainigszeiten" +
		"\n\t\t\t\u0060!Training <value>\u0060: Aktualisiert Trainingszeiten (z.B. *!Training \"our **new** trainings are ...\"*). Fett oder Kursive Schrift? Check out Discord Markup\n\t\t\t:arrow_right:" + DiscordMarkupHelpURL +
		"\n\t\t\t\u0060!Stats <battletag>\u0060: Spieler Statistiken, der Spieler sollte zuvor mit *!Register* registriert werden. (z.B. *!Stats Krusher-9911*)" +
		"\n\t\t\t\u0060!Register <battletag>\u0060: Registriert neuen Spieler (z.B. *!Register Krusher-9911*)" +
		"\n\t\t\t\u0060!Update <battletag>\u0060: Aktualisiert Statistik für angegebenen Spieler (Statistiken werden täglich einmal automatisch aktualisiert) (z.B. *!Update Krusher-9911*)" +
		"\n\t\t\t`!Config <platform=value region=value>`: Erstellt Konfiguration für Platform und Region des Overwatch Teams um Statistiken zu nutzen." +
		" Standard ist pc und eu. Creates configuration for platform and region for the overwatch team. Possible values for platform = pc,psn (PlayStation),xbl(Xbox) | region=eu,us,asia. **Note** if your platform is pc you must specify also the region. If your platform is psn or xbl you need to only specify platform." +
		"e.g.!Config platform=pc region=eu or platform=psn"

	//CommandTraining = "**Trainings**:\r\nMontag: ab 19:30 (Scrim, Review)\r\nDienstag: ab 19:30 (Scrim, Review)\r\nDonnerstag ab 19:30 (Ranked)"
	PlatformPC   = "pc"
	PlatformPS   = "psn"
	PlatformXbox = "xbl"

	RegionEU   = "eu"
	RegionUS   = "us"
	RegionAsia = "asia"
)

var (
	commandMap = map[string]getCommandContent{
		"!Training": getTrainingTimes,
		"!Help":     getCurrentlySupportedCommands,
		"!Stats":    getOverwatchPlayerStats,
		"!Register": setNewOverwatchPlayer,
		"!Update":   setNewOverwatchPlayer,
		"!Config":   setGuildConfig,
	}

	platforms = []string{PlatformPC, PlatformPS, PlatformXbox}
	regions   = []string{RegionEU, RegionUS, RegionAsia}
)

type getCommandContent func(params []string) string

func verfiyPlatform(val string) bool {

	return utils.ContainsString(platforms, val)
}

func verifyRegion(val string) bool {

	return utils.ContainsString(regions, val)
}

func setGuildConfig(params []string) string {
	if params == nil {
		return fmt.Sprintf("You need at least one of the following setting parameters. region=eu and/or platform=pc. !Help for further information.")
	}

	var platform string
	var region string
	for _, param := range params {
		paramStruct := strings.Split(param, "=")
		switch paramStruct[0] {
		case "platform":
			if verfiyPlatform(paramStruct[1]) {
				platform = paramStruct[1]
			} else {
				return fmt.Sprintf("Your defined platform is not valid. It must be pc,psn (PlayStation) or xbl(Xbox). !Help for further information.")
			}
		case "region":
			if verifyRegion(paramStruct[1]) {
				region = paramStruct[1]
			} else {
				return fmt.Sprintf("Your defined region is not valid. It must be eu, us or asia. !Help for further information.")
			}
		}
	}

	if platform == PlatformPS || platform == PlatformXbox {
		region = ""
	}
	if platform == PlatformPC && region == "" {
		return fmt.Sprintf("If you define pc as platform you need also define your region (eu,us,asia). !Help for further information.")
	}

	guildSettings := guildSettingsPersistenceLayer{Platform: platform, Region: region}
	if err := thisSession.db.setGuildConfig(thisSession.ws.cachedMessagePayload.GuildId, &guildSettings); err != nil {
		return fmt.Sprintf("Error while writing guild config.")
	}

	return fmt.Sprintf("Created/Updated config. *%v* *%v*", platform, region)
}

//noinspection GoUnusedParameter
func getTrainingTimes(params []string) string {
	//Save param as new Training Content in DB
	if params != nil {
		if err := thisSession.db.updateTrainingDates(thisSession.ws.cachedMessagePayload.GuildId, trainingDatesPersistenceLayer{params[len(params)-1]}); err != nil {
			return fmt.Sprintf("Error updating Training dates: **%v**\n*%v*\n", params[len(params)-1], string(err.Error()))
		}
		return fmt.Sprintf("Training dates set to:\n*%v*", params[len(params)-1])
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
	param := strings.Replace(params[0], "#", "-", 1)

	var config guildSettingsPersistenceLayer
	if err := thisSession.db.getGuildConfig(thisSession.ws.cachedMessagePayload.GuildId, &config); err != nil {
		//Take default if guild config doesnt exist not existing
	}

	//Set defaults
	if config.Platform == "" {
		config.Platform = "pc"
		config.Region = "eu"
	}

	owPlayerLiveStats, err := getPlayerStats(param, config.Platform, config.Region)
	if err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}
	var owPlayerPersistenceStats owStatsPersistenceLayer
	var info string
	info = "Tip: If you want the stats for your training session instead of the whole day you need to call !Update before your training."
	if err = thisSession.db.readPlayer(param, &owPlayerPersistenceStats); err != nil {
		info = fmt.Sprintf("The requested player is not registered therefore the statistics containing the data of the whole current season. If you want your global and daily statistics you need to call `!Register %v` first.", param)
	}

	return fmt.Sprintf(":chart_with_upwards_trend:Statistik für Spieler: **%v**\nRating: **%v**\nCompetitive Games played (all): *%v* Games won (all): *%v*\nTrend: *%d*sr (started today at *%v*)\nGames played today: *%v*\nGames won today: *%v*\n**%v**",
		owPlayerLiveStats.Name,
		owPlayerLiveStats.Rating,
		owPlayerLiveStats.CompetitiveStats.Games.Played,
		owPlayerLiveStats.CompetitiveStats.Games.Won,
		owPlayerLiveStats.Rating-owPlayerPersistenceStats.OWPlayer.Rating,
		owPlayerPersistenceStats.OWPlayer.Rating,
		owPlayerLiveStats.CompetitiveStats.Games.Played-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Played,
		owPlayerLiveStats.CompetitiveStats.Games.Won-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Won,
		info,
	)
}

func setNewOverwatchPlayer(params []string) string {
	param := strings.Replace(params[0], "#", "-", 1)

	var config guildSettingsPersistenceLayer
	if err := thisSession.db.getGuildConfig(thisSession.ws.cachedMessagePayload.GuildId, &config); err != nil {
		//Take default if guild config doesnt exist not existing
	}

	//Set defaults
	if config.Platform == "" {
		config.Platform = "pc"
		config.Region = "eu"
	}

	owPlayerLiveStats, err := getPlayerStats(param, config.Platform, config.Region)
	if err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}
	owStatsPersistenceLayer := owStatsPersistenceLayer{Battletag: param, OWPlayer: *owPlayerLiveStats, Guild: thisSession.ws.cachedMessagePayload.GuildId}
	if err = thisSession.db.writePlayer(owStatsPersistenceLayer); err != nil {
		return fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
	}
	return fmt.Sprintf("Player **%v** added/refreshed.", param)
}
