package main

import (
	"fmt"
	"github.com/revel/cmd/utils"
	"strconv"
	"strings"
)

const (
	DiscordMarkupHelpURL = "https://gist.github.com/Almeeida/41a664d8d5f3a8855591c2f1e0e07b19"
	PlatformPC   = "pc"
	PlatformPS   = "psn"
	PlatformXbox = "xbl"

	RegionEU   = "eu"
	RegionUS   = "us"
	RegionAsia = "asia"

	ErrorIcon     = "https://freeiconshop.com/wp-content/uploads/edd/error-flat.png"
	WarningIcon   = "https://www.pinclipart.com/picdir/middle/202-2022729_triangular-clipart-safety-sign-warning-icon-png-transparent.png"
	ErrorFooter   = "Please try again later. If this error remains, please contact our support by creating an issue on github: https://github.com/LukasMarckmiller/OverwatchDiscordBot/issues"
	OverwatchIcon = "http://www.stickpng.com/assets/images/586273b931349e0568ad89df.png"

	//Info Messages
	TipMarkup             = "Tip: You can pimp your text with discord Markups like bold,italic text or you can use discord Emojis with :emoji_name:. For a newline insert \\r\\n into your text."
	TipChangeTraining     = "Tip: If you want to change the training days just type !Training followed by some text (e.g. !Training \"our new dates\\r\\n\"). You can also use discords Markup for bold, italic or some other styles or emotes with :emote:. Use \\r\\n for a newline."
	InfoUnderConstruction = "Note: This bot is still under construction. Stored data can be removed, or Commands renamed any time while this bot is not official released"
	//Error Messages
	ErrorGuilNoParams          = "You need at least one of the following setting parameters. region=eu and/or platform=pc. !Help for further information."
	ErrorGuildPlatformNotValid = "Your defined platform is not valid. It must be pc,psn (PlayStation) or xbl(Xbox). !Help for further information."
	ErrorGuildRegionNotValid   = "Your defined region is not valid. It must be eu, us or asia. !Help for further information."
	ErrorGuildReqionRequired   = "If you define pc as platform you need also define your region (eu,us,asia). !Help for further information."
	//Help Messages
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

type getCommandContent func(params []string) discordMessageRequest

func verfiyPlatform(val string) bool {

	return utils.ContainsString(platforms, val)
}

func verifyRegion(val string) bool {

	return utils.ContainsString(regions, val)
}

func setGuildConfig(params []string) (discordMessageRequest discordMessageRequest) {
	if params == nil {
		return getErrorMessageRequest(ErrorGuilNoParams)
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
				return getErrorMessageRequest(ErrorGuildPlatformNotValid)
			}
		case "region":
			if verifyRegion(paramStruct[1]) {
				region = paramStruct[1]
			} else {
				return getErrorMessageRequest(ErrorGuildRegionNotValid)
			}
		}
	}

	if platform == PlatformPS || platform == PlatformXbox {
		region = ""
	}
	if platform == PlatformPC && region == "" {
		return getErrorMessageRequest(ErrorGuildReqionRequired)
	}

	guildSettings := guildSettingsPersistenceLayer{Platform: platform, Region: region}
	if err := thisSession.db.setGuildConfig(thisSession.ws.cachedMessagePayload.GuildId, &guildSettings); err != nil {
		return getErrorMessageRequest("Error while writing guild config.")
	}

	discordMessageRequest.Embed.Author.Name = "Discord Server Config Created/Updated"
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
	return
}

//noinspection GoUnusedParameter
func getTrainingTimes(params []string) (discordMessageRequest discordMessageRequest) {
	//Save param as new Training Content in DB
	if params != nil {
		if err := thisSession.db.updateTrainingDates(thisSession.ws.cachedMessagePayload.GuildId, trainingDatesPersistenceLayer{params[len(params)-1]}); err != nil {
			return getErrorMessageRequest(fmt.Sprintf("Error updating Training dates: **%v**\n*%v*\n", params[len(params)-1], string(err.Error())))
		}
		discordMessageRequest.Embed.Author.Name = "Updated Training days"
		discordMessageRequest.Embed.Description = params[len(params)-1]
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		discordMessageRequest.Embed.Footer.Text = TipMarkup
		return
	}
	var dates trainingDatesPersistenceLayer
	if err := thisSession.db.getTrainingDates(thisSession.ws.cachedMessagePayload.GuildId, &dates); err != nil {
		return getErrorMessageRequest(fmt.Sprintf("Error while retrieving training dates:\n*%v*\n", string(err.Error())))
	}

	discordMessageRequest.Embed.Author.Name = "Training Days"
	discordMessageRequest.Embed.Description = dates.Value
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
	discordMessageRequest.Embed.Footer.Text = TipChangeTraining
	return
}

//noinspection GoUnusedParameter
func getCurrentlySupportedCommands(params []string) (discordMessageRequest discordMessageRequest) {
	//param unused
	discordMessageRequest.Embed.Author.Name = "OverwatchTeam Discord Bot - Help"
	discordMessageRequest.Embed.Title = "All currently supported Commands with examples"
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
	discordMessageRequest.Embed.Footer.Text = InfoUnderConstruction
	discordMessageRequest.Embed.Footer.IconUrl = WarningIcon
	discordMessageRequest.Embed.Fields = []discordEmbedFieldObject{
		{Name: "!Training", Value: "Displays current Training days"},
		{Name: "!Training <value>", Value: "Updates Training days (e.g. *!Training \"our **new** trainings are ...\"*). Bold, Italic... Style? Check out Discord Markup:arrow_right:" + DiscordMarkupHelpURL},
		{Name: "!Stats <battletag>", Value: "Displays Player statistics. Player should be registered before *!Register* (e.g. *!Stats Krusher-9911*)"},
		{Name: "!Register <battletag>", Value: "Registers new player. Registered players statistics getting updated automatically every day. (e.g. *!Register Krusher-9911*)"},
		{Name: "!Update <battletag>", Value: "Updates players statistics and stores it or registers the player if not existing. (e.g. *!Update Krusher-9911*)"},
		{Name: "!Config <platform=value region=value>", Value: "Creates a server config with region and platform to use the Overwatch stats also for Playstation or XboxPlayers. Supported Platforms are pc, xbl (XBox) or psn (PlayStation)." +
			"Supported Regions are eu,us and asia. Note if your overwatch team is playing on XBox or Playstation, you only need to specify the platform and not the region. (e.g. *!Config platform=psn* for PlayStation or *!Config platform=pc region=us* for PC/US "},}
	return
}

func getOverwatchPlayerStats(params []string) (messageObject discordMessageRequest) {

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
		messageObject.Embed.Color = 0xff0000
		messageObject.Embed.Author.Name = "Error"
		messageObject.Embed.Description = fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
		messageObject.Embed.Thumbnail.Url = ErrorIcon
		messageObject.Embed.Footer.Text = ErrorFooter
	}
	var owPlayerPersistenceStats owStatsPersistenceLayer

	messageObject.Embed.Footer.Text = "Tip: You probably need to close and start Overwatch in order to get the newest stats. If you want the stats for your training session instead of the whole day you need to call !Update before your training."
	if err = thisSession.db.readPlayer(param, &owPlayerPersistenceStats); err != nil {
		messageObject.Embed.Footer.Text = fmt.Sprintf("The requested player is not registered therefore the statistics containing the data of the whole current season. If you want your global and daily statistics you need to call `!Register %v` first.", param)
	}

	var winrateAll int
	var winrateToday int

	if owPlayerLiveStats.CompetitiveStats.Games.Played != 0 {
		winrateAll = int(float32(owPlayerLiveStats.CompetitiveStats.Games.Won) / float32(owPlayerLiveStats.CompetitiveStats.Games.Played) * 100.0)
	}
	if owPlayerLiveStats.CompetitiveStats.Games.Played-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Played != 0 {
		winrateToday = int(float32(owPlayerLiveStats.CompetitiveStats.Games.Won-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Won) /
			float32(owPlayerLiveStats.CompetitiveStats.Games.Played-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Played) * 100.0)
	}

	messageObject.Embed.Author.Name = "Overwatch Player Statistics"
	messageObject.Embed.Author.IconUrl = owPlayerLiveStats.Icon
	messageObject.Embed.Title = owPlayerLiveStats.Name
	messageObject.Embed.Thumbnail.Url = owPlayerLiveStats.RatingIcon
	messageObject.Embed.Color = 0x970097
	messageObject.Embed.Description = "Competitive Game Mode"
	messageObject.Embed.Fields = []discordEmbedFieldObject{
		{Name: "Rating", Value: strconv.Itoa(owPlayerLiveStats.Rating) + " SR", Inline: true},
		{Name: "Trend", Value: strconv.Itoa(owPlayerLiveStats.Rating-owPlayerPersistenceStats.OWPlayer.Rating) + " SR", Inline: true},
		{Name: "Played (all)", Value: strconv.Itoa(owPlayerLiveStats.CompetitiveStats.Games.Played, ), Inline: true},
		{Name: "Won (all)", Value: fmt.Sprintf("%d  Winrate: %d%%", owPlayerLiveStats.CompetitiveStats.Games.Won, winrateAll), Inline: true},
		{Name: "Played (today)", Value: strconv.Itoa(owPlayerLiveStats.CompetitiveStats.Games.Played - owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Played), Inline: true},
		{Name: "Won (today)", Value: fmt.Sprintf("%d  Winrate: %d%%",
			owPlayerLiveStats.CompetitiveStats.Games.Won-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Won, winrateToday), Inline: true},
	}
	return
	/*fmt.Sprintf(":chart_with_upwards_trend:Statistik für Spieler: **%v**\nRating: **%v**\nCompetitive Games played (all): *%v* Games won (all): *%v*\nTrend: *%d*sr (started today at *%v*)\nGames played today: *%v*\nGames won today: *%v*\n**%v**",
	owPlayerLiveStats.Name,
	owPlayerLiveStats.Rating,
	owPlayerLiveStats.CompetitiveStats.Games.Played,
	owPlayerLiveStats.CompetitiveStats.Games.Won,
	owPlayerLiveStats.Rating-owPlayerPersistenceStats.OWPlayer.Rating,
	owPlayerPersistenceStats.OWPlayer.Rating,
	owPlayerLiveStats.CompetitiveStats.Games.Played-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Played,
	owPlayerLiveStats.CompetitiveStats.Games.Won-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Won,
	info,*/
}
func setNewOverwatchPlayer(params []string) (discordMessageRequest discordMessageRequest) {
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
		return getErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
	}
	owStatsPersistenceLayer := owStatsPersistenceLayer{Battletag: param, OWPlayer: *owPlayerLiveStats, Guild: thisSession.ws.cachedMessagePayload.GuildId}
	if err = thisSession.db.writePlayer(owStatsPersistenceLayer); err != nil {
		return getErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
	}
	discordMessageRequest.Embed.Author.Name = owPlayerLiveStats.Name
	discordMessageRequest.Embed.Author.IconUrl = owPlayerLiveStats.Icon
	discordMessageRequest.Embed.Title = "Player added/refreshed"
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
	discordMessageRequest.Embed.Footer.Text = "Tip: To track your sr for each training, just type !Update " + owPlayerLiveStats.Name + " before each training. After or during the Trainig you can see your progress with !Stats " + owPlayerLiveStats.Name
	discordMessageRequest.Content = fmt.Sprintf("Player **%v** added/refreshed.", param)
	return
}

func getErrorMessageRequest(message string) (request discordMessageRequest) {
	request.Embed.Color = 0xff0000
	request.Embed.Author.Name = "Error"
	request.Embed.Description = message
	request.Embed.Thumbnail.Url = ErrorIcon
	request.Embed.Footer.Text = ErrorFooter
	return
}
