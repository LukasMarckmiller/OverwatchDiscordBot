package main

import (
	"fmt"
	"github.com/revel/cmd/utils"
	"strconv"
	"strings"
	"time"
)

const (
	DiscordMarkupHelpURL = "https://gist.github.com/Almeeida/41a664d8d5f3a8855591c2f1e0e07b19"
	PlatformPC           = "pc"
	PlatformPS           = "psn"
	PlatformXbox         = "xbl"

	RegionEU   = "eu"
	RegionUS   = "us"
	RegionAsia = "asia"

	ErrorIcon     = "https://freeiconshop.com/wp-content/uploads/edd/error-flat.png"
	ErrorFooter   = "Please try again later. If this error remains, please contact our support by creating an issue on github: https://github.com/LukasMarckmiller/OverwatchDiscordBot/issues"
	OverwatchIcon = "http://www.stickpng.com/assets/images/586273b931349e0568ad89df.png"

	//Info Messages
	TipMarkup             = "Tip: You can pimp your text with discord Markups like bold,italic text or you can use discord Emojis with :emoji_name:. For a newline insert \\r\\n into your text."
	TipChangeTraining     = "Tip: If you want to change the training days just type !Training followed by some text (e.g. !Training \"our new dates\\r\\n\"). You can also use discords Markup for bold, italic or some other styles or emotes with :emote:. Use \\r\\n for a newline."
	TipUpdateProfile      = "Tip: You probably need to close and start Overwatch in order to get the newest stats. If you want the stats for your training session instead of the whole day you need to call !Update before your training."
	TipPollCreated        = "Tip: If you already created a poll, you can check the status with another !Poll call."
	TipPollUpdate         = "Tip: You can accept a poll with !+ or decline it with !-. Note: You have to be in the same Channel the poll started to accept or decline it!"
	TipPollAccept         = "Tip: You can specify a reason when you decline a poll with !- \"the reason comes here\"."
	TipCertainComp        = "Tip: If you want additional infos to a certain comp just type !Comps followed by the Name of the comp you are looking for. **NOTE: FEATURE UNDER CONSTRUCTION**"
	InfoPollTimeout       = "Note: A poll times out after 5 min. This time cant be changed by the user."
	InfoUnderConstruction = "Note: This bot is still under construction. Stored data can be removed, or Commands renamed any time while this bot is not official released."
	//Error Messages
	ErrorGuildNoParams         = "You need at least one of the following setting parameters. region=eu and/or platform=pc. !Help for further information."
	ErrorGuildPlatformNotValid = "Your defined platform is not valid. It must be pc,psn (PlayStation) or xbl(Xbox). !Help for further information."
	ErrorGuildRegionNotValid   = "Your defined region is not valid. It must be eu, us or asia. !Help for further information."
	ErrorGuildReqionRequired   = "If you define pc as platform you need also define your region (eu,us,asia). !Help for further information."
	//Help Messages

	Timeout = 10 * time.Minute

	//Overwatch role icon ids
	Tank    = "<:tank:580725264435380224>"
	Support = "<:heal:580725264422535184>"
	Dps     = "<:dps:580725264322002954>"

	//Overwatch hero icon ids

	Zarya        = "<:zarya:580725388276269076>"
	Reinhardt    = "<:reinhardt:580725264250568745>"
	Mcree        = "<:mcree:580725264087253012>"
	Widowmaker   = "<:widowmaker:580725264154230795>"
	Lucio        = "<:lucio:580725263994716190>"
	Mercy        = "<:mercy:580725264263413760>"
	Winston      = "<:winston:580725264405889057>"
	Dva          = "<:dva:580725264082796544>"
	Tracer       = "<:tracer:580725264351494172>"
	Genji        = "<:genji:580725263696920598>"
	Zen          = "<:zen:580725264762404879>"
	Doomfist     = "<:doomfist:580729098348003328>"
	Bastion      = "<:bastion:580729763832922113>"
	Roadhog      = "<:roadhog:580730509970243584>"
	Hanzo        = "<:hanzo:580731268170514433>"
	Ana          = "<:ana:580731268833083402>"
	Soldier      = "<:soldier:580731909978718258>"
	Mei          = "<:mei:580732628903264256>"
	Reaper       = "<:reaper:580732628794212390>"
	Pharah       = "<:pharah:580737542190792704>"
	Wreckingball = "<:wreckingball:580737541976752138>"
	Sombra       = "<:sombra:580737542236930058>"
	Brigitte     = "<:brigitte:580737541716574233>"
	Moira        = "<:moira:580737542094061571>"
	Junkrat      = "<:junkrat:580737542186467329>"
	Orisa        = "<:orisa:580737541821693963>"
	Baptiste     = "<:baptiste:580738605459308544>"
	Torbjrn      = "<:torbjrn:580738778017169421>"
)

var (
	commandMap = map[string]getCommandContent{
		"training":   getTrainingTimes,
		"help":       getCurrentlySupportedCommands,
		"stats":      getOverwatchPlayerStats,
		"register":   setNewOverwatchPlayer,
		"update":     setNewOverwatchPlayer,
		"config":     setGuildConfig,
		"poll":       startReadyPoll,
		"+":          setUserReady,
		"-":          setUserNotReady,
		"deletePoll": removePoll,
		"comps":      getAllCompositions,
	}

	platforms = []string{PlatformPC, PlatformPS, PlatformXbox}
	regions   = []string{RegionEU, RegionUS, RegionAsia}

	pollCache = map[string]pollCacheObject{}
)

type getCommandContent func(params []string)

func getAllCompositions(params []string) {
	//params currently unused
	var firstPage discordMessageRequest
	firstPage.Embed.Author.Name = "Overwatch All Compositions 2016-2018"
	firstPage.Embed.Title = "Composition Dictionary 1/2"
	firstPage.Embed.Description = "A historical overview of meta compositions in Overwatch for intermediate viewers."
	firstPage.Embed.Color = 0x970097
	firstPage.Embed.Thumbnail.Url = OverwatchIcon
	firstPage.Embed.Footer.Text = TipCertainComp

	//First page
	firstPage.Embed.Fields = []discordEmbedFieldObject{
		{Name: "Classic Death Ball", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Zarya, Dps, Mcree, Widowmaker, Support, Mercy, Lucio)},
		{Name: "Classic Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Winston, Dva, Dps, Genji, Tracer, Support, Zen, Lucio)},
		{Name: "El Presidente", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Dva, Dps, Bastion, Mcree, Support, Lucio, Mercy)},
		{Name: "Double Sniper", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Dva, Dps, Hanzo, Widowmaker, Support, Ana, Lucio)},
		{Name: "The 2-3-1", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Winston, Dva, Dps, Tracer, Genji, Soldier, Support, Lucio)},
		{Name: "Classic Anti-dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Dva, Roadhog, Dps, Mcree, Support, Ana, Lucio)},
		{Name: "The \"N.I.P\" Triple Tank", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Zarya, Roadhog, Dps, Soldier, Support, Ana, Lucio)},
		{Name: "Beyblade/Mei-Reaper", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Zarya, Dps, Mei, Reaper, Support, Ana, Lucio)},
		{Name: "Triple DPS Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Winston, Dps, Tracer, Genji, Soldier, Support, Zen, Lucio)},
		{Name: "The \"EnVyUs\" Triple Tank", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Dva, Roadhog, Dps, Soldier, Support, Ana, Lucio)},
		{Name: "Nanovisor/Nanoblade", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Winston, Dva, Dps, Genji, Soldier, Support, Ana, Lucio)},
		{Name: "Nanovisor/Nanoblade", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s", Tank, Reinhardt, Zarya, Dva, Roadhog, Support, Ana, Lucio)},
		{Name: "Sombra as Support (No longer possible)", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Winston, Dva, Dps, Tracer, Soldier, Support, Lucio, Sombra)},
		{Name: "The \"Selfless\"", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Roadhog, Dps, Tracer, Soldier, Support, Ana, Lucio)},
		{Name: "Pharah Mercy Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Winston, Dva, Dps, Tracer, Pharah, Support, Mercy, Lucio)},
		{Name: "New Anti-Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Zarya, Dps, Reaper, Junkrat, Support, Ana, Lucio)},
	}
	_, _ = sendMessage(firstPage)
	//Second page
	var secondPage discordMessageRequest
	secondPage.Embed.Author.Name = "Overwatch All Compositions 2018-2019"
	secondPage.Embed.Title = "Composition Dictionary 2/2"
	secondPage.Embed.Description = "A historical overview of meta compositions in Overwatch for intermediate viewers."
	secondPage.Embed.Color = 0x970097
	secondPage.Embed.Thumbnail.Url = OverwatchIcon
	secondPage.Embed.Footer.Text = TipCertainComp
	secondPage.Embed.Fields = []discordEmbedFieldObject{
		{Name: "Doomfist-McCree", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Reinhardt, Zarya, Dps, Doomfist, Mcree, Support, Ana, Lucio)},
		{Name: "Junkrat-Widow Defense", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Orisa, Dva, Dps, Junkrat, Widowmaker, Support, Zen, Mercy)},
		{Name: "Orisa-Hog", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Orisa, Roadhog, Dps, Hanzo, Widowmaker, Support, Zen, Mercy)},
		{Name: "Hog-Dva", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Roadhog, Dva, Dps, Tracer, Widowmaker, Support, Zen, Mercy)},
		{Name: "Pirate Ship", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Orisa, Roadhog, Dps, Bastion, Hanzo, Support, Zen, Mercy)},
		{Name: "Tracer-Sombra Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Winston, Dva, Dps, Tracer, Sombra, Support, Moira, Lucio)},
		{Name: "GOATS 1.0", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s", Tank, Reinhardt, Zarya, Dva, Support, Brigitte, Moira, Lucio)},
		{Name: "Disruptive Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Wreckingball, Dva, Dps, Doomfist, Sombra, Support, Ana, Lucio)},
		{Name: "GOATS 2.0", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s", Tank, Reinhardt, Zarya, Dva, Support, Brigitte, Zen, Lucio)},
		{Name: "Triple Support Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s ", Tank, Winston, Dva, Dps, Sombra, Support, Ana, Lucio, Brigitte)},
		{Name: "The \"Chengdu Hunters\"", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s ", Tank, Wreckingball, Dva, Dps, Sombra, Pharah, Support, Zen, Mercy)},
		{Name: "Bunker", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Orisa, Roadhog, Dps, Bastion, Hanzo, Support, Baptiste, Mercy)},
		{Name: "Ball Dive(Anti Bunker)", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Wreckingball, Dva, Dps, Genji, Sombra, Support, Ana, Lucio)},
		{Name: "Gold Elo Classic", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", Tank, Zarya, Dps, Genji, Tracer, Torbjrn, Widowmaker, Support, Moira)},
	}
	_, _ = sendMessage(secondPage)
	return
}

//params currently unused
func removePoll(params []string) {
	var discordMessageRequest discordMessageRequest
	guildId := thisSession.ws.cachedMessagePayload.GuildId
	channelId := thisSession.ws.cachedMessagePayload.ChannelId
	cachedPoll, ok := pollCache[guildId+channelId]
	if ok {
		delete(pollCache, cachedPoll.Guild+cachedPoll.Channel)
		avatarUrl, _ := thisSession.ws.getUserAvatarOrDefaultUrl(cachedPoll.Creator.Id, cachedPoll.Creator.Avatar, cachedPoll.Creator.Discriminator)
		discordMessageRequest.Embed.Author.Name = cachedPoll.Creator.Username + "#" + cachedPoll.Creator.Discriminator + "´s Poll"
		discordMessageRequest.Embed.Author.IconUrl = avatarUrl
		discordMessageRequest.Embed.Title = "Poll successfully deleted."
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		discordMessageRequest.Embed.Footer.Text = TipPollCreated
		_, _ = sendMessage(discordMessageRequest)
		return
	} else {
		sendInfoMessageRequest("No Poll was created for this chanel to delete. First create a Poll with !Poll <n>. !Help for more details.")
		return
	}
}

func countReadyMembers(cachedObject pollCacheObject) (count int) {
	for _, member := range cachedObject.Members {
		if member.Ready {
			//If someone is not ready wait for everybody to be ready
			count++
		}
	}
	return
}

func checkIfPollIsDone(cachedPollObject pollCacheObject) {

	//Everybody is ready
	if len(cachedPollObject.Members) == cachedPollObject.Size {
		//Check if everybody responed with ready
		for _, member := range cachedPollObject.Members {
			if !member.Ready {
				//If someone is not ready wait for everybody to be ready
				return
			}
		}

		delete(pollCache, cachedPollObject.Guild+cachedPollObject.Channel)

		avatarUrl, _ := thisSession.ws.getUserAvatarOrDefaultUrl(cachedPollObject.Creator.Id, cachedPollObject.Creator.Avatar, cachedPollObject.Creator.Discriminator)

		discordMessageRequest := discordMessageRequest{}
		discordMessageRequest.Embed.Author.Name = cachedPollObject.Creator.Username + "#" + cachedPollObject.Creator.Discriminator + "´s Poll"
		discordMessageRequest.Embed.Author.IconUrl = avatarUrl
		discordMessageRequest.Embed.Title = "Poll is finished everybody has responded and is ready."
		discordMessageRequest.Embed.Description = fmt.Sprintf("%d out of %d are ready to go!", countReadyMembers(cachedPollObject), cachedPollObject.Size)
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		discordMessageRequest.Embed.Footer.Text = InfoPollTimeout
		discordMessageRequest.Content = "<@" + cachedPollObject.Creator.Id + ">"
		var cachedPollMembers []discordEmbedFieldObject
		for _, val := range cachedPollObject.Members {
			cachedPollMembers = append(cachedPollMembers, discordEmbedFieldObject{Name: val.User.Username, Value: getReadyStatValue(val.Ready, val.Reason)})
		}
		discordMessageRequest.Embed.Fields = cachedPollMembers
		_, _ = sendMessage(discordMessageRequest)
	}
}

func existsUserInPollCache(cachedPoll pollCacheObject, user discordUserObject) (int, bool) {
	for index, member := range cachedPoll.Members {
		if member.User.Id == user.Id {
			return index, true
		}
	}

	return -1, false
}

func setUserNotReady(params []string) {
	var discordMessageRequest discordMessageRequest
	guildId := thisSession.ws.cachedMessagePayload.GuildId
	channelId := thisSession.ws.cachedMessagePayload.ChannelId
	var response string
	if params != nil {
		response = params[0]
	}
	cachedPoll, ok := pollCache[guildId+channelId]
	if ok {
		cachedAuthor := thisSession.ws.cachedMessagePayload.Author
		if userIndex, exists := existsUserInPollCache(cachedPoll, cachedAuthor); exists {
			cachedPoll.Members[userIndex] = readyCheckMember{User: cachedAuthor, Reason: response, Ready: false}
		} else {
			if len(cachedPoll.Members) == cachedPoll.Size {
				sendInfoMessageRequest("Poll already full of members. Make a bigger Poll next time!")
				return
			}
			cachedPoll.Members = append(cachedPoll.Members, readyCheckMember{User: cachedAuthor, Ready: false, Reason: response})
		}
		//reassign to update value
		avatarUrl, _ := thisSession.ws.getUserAvatarOrDefaultUrl(cachedPoll.Creator.Id, cachedPoll.Creator.Avatar, cachedPoll.Creator.Discriminator)

		pollCache[guildId+channelId] = cachedPoll
		discordMessageRequest.Embed.Author.Name = cachedAuthor.Username + " is not ready "
		discordMessageRequest.Embed.Author.IconUrl = avatarUrl
		discordMessageRequest.Embed.Description = fmt.Sprintf("Poll times out in %.0fmin", cachedPoll.CreationTime.Add(Timeout).Sub(time.Now()).Minutes())
		discordMessageRequest.Embed.Title = response + " :expressionless:"
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		discordMessageRequest.Embed.Footer.Text = TipPollAccept
		_, _ = sendMessage(discordMessageRequest)
		checkIfPollIsDone(cachedPoll)
		return
	} else {
		sendInfoMessageRequest("You have to create a poll first in order to accept or decline. Polls can be created with !Poll <num of members>.")
		return
	}
}

func setUserReady(params []string) {
	var discordMessageRequest discordMessageRequest
	guildId := thisSession.ws.cachedMessagePayload.GuildId
	channelId := thisSession.ws.cachedMessagePayload.ChannelId

	cachedPoll, ok := pollCache[guildId+channelId]
	if ok {
		cachedAuthor := thisSession.ws.cachedMessagePayload.Author
		//get poll member if already in list
		if userIndex, exists := existsUserInPollCache(cachedPoll, cachedAuthor); exists {
			cachedPoll.Members[userIndex].Ready = true
		} else {
			if len(cachedPoll.Members) == cachedPoll.Size {
				sendInfoMessageRequest("Poll already full of members. Make a bigger Poll next time!")
				return
			}
			cachedPoll.Members = append(cachedPoll.Members, readyCheckMember{User: cachedAuthor, Ready: true})
		}
		avatarUrl, _ := thisSession.ws.getUserAvatarOrDefaultUrl(cachedPoll.Creator.Id, cachedPoll.Creator.Avatar, cachedPoll.Creator.Discriminator)

		//reassign to update value
		pollCache[guildId+channelId] = cachedPoll
		discordMessageRequest.Embed.Author.Name = cachedAuthor.Username + " is ready now!"
		discordMessageRequest.Embed.Author.IconUrl = avatarUrl
		discordMessageRequest.Embed.Description = fmt.Sprintf("Poll times out in %.0fmin", cachedPoll.CreationTime.Add(Timeout).Sub(time.Now()).Minutes())
		discordMessageRequest.Embed.Title = ":ok_hand:"
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		discordMessageRequest.Embed.Footer.Text = TipPollAccept
		_, _ = sendMessage(discordMessageRequest)
		checkIfPollIsDone(cachedPoll)
		return
	} else {
		sendInfoMessageRequest("You have to create a poll first in order to accept or decline. Polls can be created with !Poll <num of members>.")
		return
	}
}

func getReadyStatValue(ready bool, reason string) string {
	if ready {
		return ":white_check_mark: Ready"
	} else {
		return ":x: Not Ready " + reason
	}
}

func startReadyPoll(params []string) {
	var discordMessageRequest discordMessageRequest
	guildId := thisSession.ws.cachedMessagePayload.GuildId
	channelId := thisSession.ws.cachedMessagePayload.ChannelId

	cachedPoll, ok := pollCache[guildId+channelId]

	if ok {
		//Existing poll
		//ignore param
		avatarUrl, _ := thisSession.ws.getUserAvatarOrDefaultUrl(cachedPoll.Creator.Id, cachedPoll.Creator.Avatar, cachedPoll.Creator.Discriminator)
		discordMessageRequest.Embed.Author.Name = cachedPoll.Creator.Username + "#" + cachedPoll.Creator.Discriminator + "´s Poll"
		discordMessageRequest.Embed.Author.IconUrl = avatarUrl
		discordMessageRequest.Embed.Title = fmt.Sprintf("Poll times out in %.0fmin", cachedPoll.CreationTime.Add(Timeout).Sub(time.Now()).Minutes())
		discordMessageRequest.Embed.Description = fmt.Sprintf("%d out of %d are ready to go!", countReadyMembers(cachedPoll), cachedPoll.Size)
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		discordMessageRequest.Embed.Footer.Text = TipPollUpdate

		var cachedPollMembers []discordEmbedFieldObject
		for _, val := range cachedPoll.Members {
			cachedPollMembers = append(cachedPollMembers, discordEmbedFieldObject{Name: val.User.Username, Value: getReadyStatValue(val.Ready, val.Reason)})
		}
		discordMessageRequest.Embed.Fields = cachedPollMembers
		_, _ = sendMessage(discordMessageRequest)
		return

	} else if params != nil { //new poll with param
		n, err := strconv.Atoi(params[0])
		if err != nil {
			sendInfoMessageRequest("The size of the poll needs to be a valid number (e.g. !Poll 5)")
			return
		}
		if n < 0 {
			sendErrorMessageRequest("Dont try to break the bot!")
			return
		} else if n == 0 {
			sendErrorMessageRequest("Dude you need at least one member for your poll.")
			return
		}

		now := time.Now()
		var pollCacheObject pollCacheObject
		pollCacheObject.Creator = thisSession.ws.cachedMessagePayload.Author
		pollCacheObject.Guild = guildId
		pollCacheObject.Channel = channelId
		pollCacheObject.CreationTime = now
		pollCacheObject.Size = n
		pollCacheObject.Members = []readyCheckMember{}
		pollCache[pollCacheObject.Guild+pollCacheObject.Channel] = pollCacheObject

		avatarUrl, _ := thisSession.ws.getUserAvatarOrDefaultUrl(pollCacheObject.Creator.Id, pollCacheObject.Creator.Avatar, pollCacheObject.Creator.Discriminator)
		discordMessageRequest.Embed.Author.Name = pollCacheObject.Creator.Username + "#" + pollCacheObject.Creator.Discriminator + " just started a new ready poll!"
		discordMessageRequest.Embed.Author.IconUrl = avatarUrl
		discordMessageRequest.Embed.Title = fmt.Sprintf("%d People involved", n)
		discordMessageRequest.Embed.Description = fmt.Sprintf("Accept with !+ or decline with !-. Poll times out in %v", Timeout)
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		discordMessageRequest.Embed.Footer.Text = TipPollCreated
		startTimer(Timeout, func() {

			cachedPoll, ok := pollCache[guildId+channelId]
			//Check if it is this poll and dont delete a poll which is created after the current poll
			if ok && now.Equal(cachedPoll.CreationTime) {
				delete(pollCache, cachedPoll.Guild+cachedPoll.Channel)
			}
		})
		_, _ = sendMessage(discordMessageRequest)
		return

	} else { //new poll but no param
		sendInfoMessageRequest("You need to specify the size of the poll (e.g. !Poll 5).")
		return
	}
}

func verfiyPlatform(val string) bool {

	return utils.ContainsString(platforms, val)
}

func verifyRegion(val string) bool {

	return utils.ContainsString(regions, val)
}

func setGuildConfig(params []string) {
	var discordMessageRequest discordMessageRequest
	if params == nil {
		sendInfoMessageRequest(ErrorGuildNoParams)
		return
	}

	var guildSettings guildSettingsPersistenceLayer

	var platform string
	var region string
	var prefix string
	for _, param := range params {
		paramStruct := strings.Split(param, "=")
		switch paramStruct[0] {
		case "platform":
			if verfiyPlatform(paramStruct[1]) {
				platform = paramStruct[1]
			} else {
				sendErrorMessageRequest(ErrorGuildPlatformNotValid)
				return
			}
		case "region":
			if verifyRegion(paramStruct[1]) {
				region = paramStruct[1]
			} else {
				sendErrorMessageRequest(ErrorGuildRegionNotValid)
				return
			}
		case "prefix":
			if paramStruct[1] == "" {
				sendErrorMessageRequest("Prefix cant be empty. Pls specify a valid prefix.")
				return
			}
			prefix = paramStruct[1]
		default:
			sendErrorMessageRequest("Unknown parameter " + paramStruct[0])
			return
		}
	}

	if platform == PlatformPS || platform == PlatformXbox {
		region = ""
	}
	if platform == PlatformPC && region == "" {
		sendErrorMessageRequest(ErrorGuildReqionRequired)
		return
	}

	//Try load settings
	if err := thisSession.db.getGuildConfig(thisSession.ws.cachedMessagePayload.GuildId, &guildSettings); err != nil {
		guildSettings = guildSettingsPersistenceLayer{Platform: platform, Region: region, Prefix: prefix}
		if err := thisSession.db.setGuildConfig(thisSession.ws.cachedMessagePayload.GuildId, &guildSettings); err != nil {
			sendErrorMessageRequest("Error while writing guild config.")
			return
		}
		discordMessageRequest.Embed.Author.Name = "Discord Server Config Created"
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		_, _ = sendMessage(discordMessageRequest)
		return
	} else { //Create new if not found
		if prefix != "" {
			guildSettings.Prefix = prefix
		}

		if platform != "" {
			guildSettings.Platform = platform
			guildSettings.Region = region
		}

		if err := thisSession.db.setGuildConfig(thisSession.ws.cachedMessagePayload.GuildId, &guildSettings); err != nil {
			sendErrorMessageRequest("Error while writing guild config.")
			return
		}

		discordMessageRequest.Embed.Author.Name = "Discord Server Config Updated"
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		_, _ = sendMessage(discordMessageRequest)
		return
	}
}

//noinspection GoUnusedParameter
func getTrainingTimes(params []string) {
	var discordMessageRequest discordMessageRequest
	//Save param as new Training Content in DB
	if params != nil {
		if err := thisSession.db.updateTrainingDates(thisSession.ws.cachedMessagePayload.GuildId, trainingDatesPersistenceLayer{params[0]}); err != nil {
			sendErrorMessageRequest(fmt.Sprintf("Error updating Training dates: **%v**\n*%v*\n", params[0], string(err.Error())))
			return
		}
		discordMessageRequest.Embed.Author.Name = "Updated Training days"
		discordMessageRequest.Embed.Description = params[0]
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
		discordMessageRequest.Embed.Footer.Text = TipMarkup
		_, _ = sendMessage(discordMessageRequest)
		return
	}
	var dates trainingDatesPersistenceLayer
	if err := thisSession.db.getTrainingDates(thisSession.ws.cachedMessagePayload.GuildId, &dates); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error while retrieving training dates:\n*%v*\n", string(err.Error())))
		return
	}

	discordMessageRequest.Embed.Author.Name = "Training Days"
	discordMessageRequest.Embed.Description = dates.Value
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
	discordMessageRequest.Embed.Footer.Text = TipChangeTraining
	_, _ = sendMessage(discordMessageRequest)
	return
}

//noinspection GoUnusedParameter
func getCurrentlySupportedCommands(params []string) {
	//param unused
	var discordMessageRequest discordMessageRequest
	config := getGuildConfigSave(thisSession.ws.cachedMessagePayload.GuildId)
	discordMessageRequest.Embed.Author.Name = "OverwatchTeam Discord Bot - Help"
	discordMessageRequest.Embed.Title = "All currently supported Commands with examples."
	discordMessageRequest.Embed.Description = "If your using Overwatch related commands make sure your profile is set to public"
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
	discordMessageRequest.Embed.Footer.Text = InfoUnderConstruction
	discordMessageRequest.Embed.Footer.IconUrl = OverwatchIcon
	discordMessageRequest.Embed.Fields = []discordEmbedFieldObject{
		{Name: config.Prefix + "Training", Value: "Displays current Training days"},
		{Name: config.Prefix + "Training <value>", Value: "Updates Training days (e.g. *!Training \"our **new** trainings are ...\"*). Bold, Italic... Style? Check out Discord Markup:arrow_right:" + DiscordMarkupHelpURL},
		{Name: config.Prefix + "Stats <battletag>", Value: "Displays Player statistics. Player should be registered before *!Register* (e.g. *!Stats Krusher-9911*)"},
		{Name: config.Prefix + "Register <battletag>", Value: "Registers new player. Registered players statistics getting updated automatically every day. (e.g. *!Register Krusher-9911*)"},
		{Name: config.Prefix + "Update <battletag>", Value: "Updates players statistics and stores it or registers the player if not existing. (e.g. *!Update Krusher-9911*)"},
		{Name: config.Prefix + "Config <platform=value region=value> <prefix=value>", Value: "Creates a server config with region and platform to use the Overwatch stats also for Playstation or XboxPlayers. And/or You can also specify a custom prefix for the bot with prefix=value e.g. (*!Config prefix=>*). Supported Platforms are pc, xbl (XBox) or psn (PlayStation)." +
			"Supported Regions are eu,us and asia. Note if your overwatch team is playing on XBox or Playstation, you only need to specify the platform and not the region. (e.g. *!Config platform=psn* for PlayStation or *!Config platform=pc region=us* for PC/US "},
		{Name: config.Prefix + "Poll <number of participants>", Value: fmt.Sprintf("Starts a ready check poll for n players. A poll times out after %d minutes.", Timeout/time.Minute)},
		{Name: config.Prefix + "Poll", Value: "Gets the status of the current poll. If everybody is ready a message is created and the creator of the poll gets tagged."},
		{Name: config.Prefix + "+", Value: "Ready."},
		{Name: config.Prefix + "- <reason>", Value: "Not ready. A reason can be passed with the command. (e.g. !- \"need water! Back in 5\"). **Note if your reason is longer then one word you need to put it in \"\"!**"},
		{Name: config.Prefix + "DeletePoll", Value: "Deletes the current Poll."},
		{Name: config.Prefix + "Comps", Value: "Returns all meta comps of Overwatch over the years."},
	}
	_, _ = sendMessage(discordMessageRequest)
	return
}

func getOverwatchPlayerStats(params []string) {
	var messageObject discordMessageRequest
	param := strings.Replace(params[0], "#", "-", 1)

	config := getGuildConfigSave(thisSession.ws.cachedMessagePayload.GuildId)

	owPlayerLiveStats, err := getPlayerStats(param, config.Platform, config.Region)
	if err != nil {
		messageObject.Embed.Color = 0xff0000
		messageObject.Embed.Author.Name = "Error"
		messageObject.Embed.Description = fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error()))
		messageObject.Embed.Thumbnail.Url = ErrorIcon
		messageObject.Embed.Footer.Text = ErrorFooter
		_, _ = sendMessage(messageObject)
		return
	}
	var owPlayerPersistenceStats owStatsPersistenceLayer

	messageObject.Embed.Footer.Text = TipUpdateProfile
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
		{Name: "Played (all)", Value: strconv.Itoa(owPlayerLiveStats.CompetitiveStats.Games.Played), Inline: true},
		{Name: "Won (all)", Value: fmt.Sprintf("%d  Winrate: %d%%", owPlayerLiveStats.CompetitiveStats.Games.Won, winrateAll), Inline: true},
		{Name: "Played (today)", Value: strconv.Itoa(owPlayerLiveStats.CompetitiveStats.Games.Played - owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Played), Inline: true},
		{Name: "Won (today)", Value: fmt.Sprintf("%d  Winrate: %d%%",
			owPlayerLiveStats.CompetitiveStats.Games.Won-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Won, winrateToday), Inline: true},
	}
	_, _ = sendMessage(messageObject)
	return

}
func setNewOverwatchPlayer(params []string) {
	var discordMessageRequest discordMessageRequest
	param := strings.Replace(params[0], "#", "-", 1)

	config := getGuildConfigSave(thisSession.ws.cachedMessagePayload.GuildId)

	owPlayerLiveStats, err := getPlayerStats(param, config.Platform, config.Region)
	if err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
		return
	}
	owStatsPersistenceLayer := owStatsPersistenceLayer{Battletag: param, OWPlayer: *owPlayerLiveStats, Guild: thisSession.ws.cachedMessagePayload.GuildId}
	if err = thisSession.db.writePlayer(owStatsPersistenceLayer); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
		return
	}
	discordMessageRequest.Embed.Author.Name = owPlayerLiveStats.Name
	discordMessageRequest.Embed.Author.IconUrl = owPlayerLiveStats.Icon
	discordMessageRequest.Embed.Title = "Player added/refreshed"
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = OverwatchIcon
	discordMessageRequest.Embed.Footer.Text = "Tip: To track your sr for each training, just type !Update " + owPlayerLiveStats.Name + " before each training. After or during the Trainig you can see your progress with !Stats " + owPlayerLiveStats.Name
	_, _ = sendMessage(discordMessageRequest)
	return
}

func sendMessage(message discordMessageRequest) (respMsg discordMessageResponse, err error) {
	respMsg, err = thisSession.ws.sendMessageToChannel(message, thisSession.ws.cachedMessagePayload.ChannelId)
	if err != nil {
		return discordMessageResponse{}, err
	}

	return respMsg, nil
}

func getGuildConfigSave(guildId string) guildSettingsPersistenceLayer {
	var config guildSettingsPersistenceLayer
	if err := thisSession.db.getGuildConfig(guildId, &config); err != nil {
		//Take default if guild config doesnt exist not existing
		config.Platform = "pc"
		config.Region = "eu"
		config.Prefix = "!"
	}
	//Take default if not set
	if config.Platform == "" {
		config.Platform = "pc"
		config.Region = "eu"
	}
	//Take default if not set
	if config.Prefix == "" {
		config.Prefix = "!"
	}

	return config
}

func loadPrefixOrDefault(guildId string) string {
	return getGuildConfigSave(guildId).Prefix
}

func sendErrorMessageRequest(message string) {
	var request discordMessageRequest
	request.Embed.Color = 0xff0000
	request.Embed.Author.Name = "Error"
	request.Embed.Description = message
	request.Embed.Thumbnail.Url = ErrorIcon
	request.Embed.Footer.Text = ErrorFooter
	_, _ = sendMessage(request)
}

func sendInfoMessageRequest(message string) {
	var request discordMessageRequest
	request.Embed.Color = 0x970097
	request.Embed.Author.Name = "Info"
	request.Embed.Description = message
	request.Embed.Thumbnail.Url = OverwatchIcon
	request.Embed.Footer.Text = InfoUnderConstruction
	_, _ = sendMessage(request)
}

type pollCacheObject struct {
	Guild        string             `json:"guild"`
	Channel      string             `json:"channel"`
	CreationTime time.Time          `json:"creation_time"`
	Size         int                `json:"size"`
	Creator      discordUserObject  `json:"creator"`
	Members      []readyCheckMember `json:"members"`
}

type readyCheckMember struct {
	User   discordUserObject `json:"user"`
	Ready  bool              `json:"ready"`
	Reason string            `json:"reason"`
}
