package main

import (
	"encoding/json"
	"fmt"
	"github.com/revel/cmd/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	//Link to Discord Markup helper document
	discordMarkupHelpURL = "https://gist.github.com/Almeeida/41a664d8d5f3a8855591c2f1e0e07b19"
	platformPC           = "pc"
	platformPS           = "psn"
	platformXbox         = "xbl"

	regionEU   = "eu"
	regionUS   = "us"
	regionAsia = "asia"

	errorIcon     = "https://freeiconshop.com/wp-content/uploads/edd/error-flat.png"
	errorFooter   = "Please try again later. If this error remains, please contact our support by creating an issue on github: https://github.com/LukasMarckmiller/OverwatchDiscordBot/issues"
	overwatchIcon = "http://www.stickpng.com/assets/images/586273b931349e0568ad89df.png"

	tipMarkup             = "Tip: You can pimp your text with discord Markups like bold,italic text or you can use discord Emojis with :emoji_name:. For a newline insert \\n into your text."
	tipChangeTraining     = "Tip: If you want to change the training days just type !Training followed by some text (e.g. !Training \"our new dates \\n\"). You can also use discords Markup for bold, italic or some other styles or emotes with :emote:. Use \\n for a newline."
	tipUpdateProfile      = "Tip: You probably need to close and start Overwatch in order to get the newest stats. If you want the stats for your training session instead of the whole day you need to call !Update before your training."
	tipPollCreated        = "Tip: If you already created a poll, you can check the status with another !Poll call."
	tipPollUpdate         = "Tip: You can accept a poll with !+ or decline it with !-. Note: You have to be in the same Channel the poll started to accept or decline it!"
	tipPollAccept         = "Tip: You can specify a reason when you decline a poll with !- \"the reason comes here\"."
	tipCertainComp        = "Tip: If you want additional infos to a certain comp just type !Comps followed by the Name of the comp you are looking for. **NOTE: FEATURE UNDER CONSTRUCTION**"
	infoPollTimeout       = "Note: A poll times out after 10 min. This time cant be changed by the user."
	infoUnderConstruction = "Note: This bot is still under construction. Stored data can be removed, or Commands renamed any time while this bot is not official released."
	//Error Messages
	errorGuildNoParams         = "You need at least one of the following setting parameters. region=eu and/or platform=pc. !Help for further information."
	errorGuildPlatformNotValid = "Your defined platform is not valid. It must be pc,psn (PlayStation) or xbl(Xbox). !Help for further information."
	errorGuildRegionNotValid   = "Your defined region is not valid. It must be eu, us or asia. !Help for further information."
	errorGuildReqionRequired   = "If you define pc as platform you need also define your region (eu,us,asia) e.g *region=eu*. !Help for further information."
	//Help Messages

	timeout = 10 * time.Minute

	//Overwatch role icon ids
	tank    = "<:tank:580725264435380224>"
	support = "<:heal:580725264422535184>"
	dps     = "<:dps:580725264322002954>"

	//Overwatch hero icon ids

	sigma        = "<:sigma:640843743129239582>"
	zarya        = "<:zarya:580725388276269076>"
	reinhardt    = "<:reinhardt:580725264250568745>"
	winston      = "<:winston:580725264405889057>"
	dva          = "<:dva:580725264082796544>"
	roadhog      = "<:roadhog:580730509970243584>"
	wreckingball = "<:wreckingball:580737541976752138>"
	orisa        = "<:orisa:580737541821693963>"

	ana      = "<:ana:580731268833083402>"
	lucio    = "<:lucio:580725263994716190>"
	mercy    = "<:mercy:580725264263413760>"
	zen      = "<:zen:580725264762404879>"
	brigitte = "<:brigitte:580737541716574233>"
	moira    = "<:moira:580737542094061571>"
	baptiste = "<:baptiste:580738605459308544>"

	mcree      = "<:mcree:580725264087253012>"
	widowmaker = "<:widowmaker:580725264154230795>"
	tracer     = "<:tracer:580725264351494172>"
	genji      = "<:genji:580725263696920598>"
	doomfist   = "<:doomfist:580729098348003328>"
	bastion    = "<:bastion:580729763832922113>"
	hanzo      = "<:hanzo:580731268170514433>"
	soldier    = "<:soldier:580731909978718258>"
	mei        = "<:mei:580732628903264256>"
	reaper     = "<:reaper:580732628794212390>"
	pharah     = "<:pharah:580737542190792704>"
	sombra     = "<:sombra:580737542236930058>"
	junkrat    = "<:junkrat:580737542186467329>"
	torbjrn    = "<:torbjrn:580738778017169421>"
	symmetra   = "<:symmetra:583253093089542166>"
	ashe       = "<:ashe:583253081094094851>"
)

var (
	HeroIconMap = map[string]string{
		"zenyatta":     zen,
		"ana":          ana,
		"zarya":        zarya,
		"reinhardt":    reinhardt,
		"mcree":        mcree,
		"widowmaker":   widowmaker,
		"lucio":        lucio,
		"mercy":        mercy,
		"winston":      winston,
		"dva":          dva,
		"tracer":       tracer,
		"genji":        genji,
		"doomfist":     doomfist,
		"bastion":      bastion,
		"roadhog":      roadhog,
		"hanzo":        hanzo,
		"soldier76":    soldier,
		"mei":          mei,
		"reaper":       reaper,
		"pharah":       pharah,
		"wreckingball": wreckingball,
		"sombra":       sombra,
		"brigitte":     brigitte,
		"moira":        moira,
		"junkrat":      junkrat,
		"orisa":        orisa,
		"baptiste":     baptiste,
		"torbjorn":     torbjrn,
		"ashe":         ashe,
		"symmetra":     symmetra,
		"sigma":        sigma,
	}
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
		"deletepoll": removePoll,
		"comps":      getAllCompositions,
	}

	platforms = []string{platformPC, platformPS, platformXbox}
	regions   = []string{regionEU, regionUS, regionAsia}

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
	firstPage.Embed.Thumbnail.Url = overwatchIcon
	firstPage.Embed.Footer.Text = tipCertainComp
	firstPage.Embed.Footer.IconUrl = overwatchIcon

	//First page
	firstPage.Embed.Fields = []discordEmbedFieldObject{
		{Name: "Classic Death Ball", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, zarya, dps, mcree, widowmaker, support, mercy, lucio)},
		{Name: "Classic Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, winston, dva, dps, genji, tracer, support, zen, lucio)},
		{Name: "El Presidente", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, dva, dps, bastion, mcree, support, lucio, mercy)},
		{Name: "Double Sniper", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, dva, dps, hanzo, widowmaker, support, ana, lucio)},
		{Name: "The 2-3-1", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, winston, dva, dps, tracer, genji, soldier, support, lucio)},
		{Name: "Classic Anti-dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, dva, roadhog, dps, mcree, support, ana, lucio)},
		{Name: "The \"N.I.P\" Triple Tank", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, zarya, roadhog, dps, soldier, support, ana, lucio)},
		{Name: "Beyblade/Mei-Reaper", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, zarya, dps, mei, reaper, support, ana, lucio)},
		{Name: "Triple DPS Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, winston, dps, tracer, genji, soldier, support, zen, lucio)},
		{Name: "The \"EnVyUs\" Triple Tank", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, dva, roadhog, dps, soldier, support, ana, lucio)},
		{Name: "Nanovisor/Nanoblade", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, winston, dva, dps, genji, soldier, support, ana, lucio)},
		{Name: "Nanovisor/Nanoblade", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s", tank, reinhardt, zarya, dva, roadhog, support, ana, lucio)},
		{Name: "Sombra as Support (No longer possible)", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, winston, dva, dps, tracer, soldier, support, lucio, sombra)},
		{Name: "The \"Selfless\"", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, roadhog, dps, tracer, soldier, support, ana, lucio)},
		{Name: "Pharah Mercy Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, winston, dva, dps, tracer, pharah, support, mercy, lucio)},
		{Name: "New Anti-Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, zarya, dps, reaper, junkrat, support, ana, lucio)},
	}
	if _, err := sendMessage(firstPage); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
		return
	}
	//Second page
	var secondPage discordMessageRequest
	secondPage.Embed.Author.Name = "Overwatch All Compositions 2018-2019"
	secondPage.Embed.Title = "Composition Dictionary 2/2"
	secondPage.Embed.Description = "A historical overview of meta compositions in Overwatch for intermediate viewers."
	secondPage.Embed.Color = 0x970097
	secondPage.Embed.Thumbnail.Url = overwatchIcon
	secondPage.Embed.Footer.Text = tipCertainComp
	secondPage.Embed.Footer.IconUrl = overwatchIcon
	secondPage.Embed.Fields = []discordEmbedFieldObject{
		{Name: "Doomfist-McCree", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, reinhardt, zarya, dps, doomfist, mcree, support, ana, lucio)},
		{Name: "Junkrat-Widow Defense", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, orisa, dva, dps, junkrat, widowmaker, support, zen, mercy)},
		{Name: "Orisa-Hog", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, orisa, roadhog, dps, hanzo, widowmaker, support, zen, mercy)},
		{Name: "Hog-Dva", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, roadhog, dva, dps, tracer, widowmaker, support, zen, mercy)},
		{Name: "Pirate Ship", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, orisa, roadhog, dps, bastion, hanzo, support, zen, mercy)},
		{Name: "Tracer-Sombra Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, winston, dva, dps, tracer, sombra, support, moira, lucio)},
		{Name: "GOATS 1.0", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s", tank, reinhardt, zarya, dva, support, brigitte, moira, lucio)},
		{Name: "Disruptive Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, wreckingball, dva, dps, doomfist, sombra, support, ana, lucio)},
		{Name: "GOATS 2.0", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s", tank, reinhardt, zarya, dva, support, brigitte, zen, lucio)},
		{Name: "Triple Support Dive", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s ", tank, winston, dva, dps, sombra, support, ana, lucio, brigitte)},
		{Name: "The \"Chengdu Hunters\"", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s ", tank, wreckingball, dva, dps, sombra, pharah, support, zen, mercy)},
		{Name: "Bunker", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, orisa, roadhog, dps, bastion, hanzo, support, baptiste, mercy)},
		{Name: "Ball Dive(Anti Bunker)", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, wreckingball, dva, dps, genji, sombra, support, ana, lucio)},
		{Name: "3-Dps (DPS variable)", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, wreckingball, dps, widowmaker, genji, ashe, support, mercy, lucio)},
		{Name: "Double Shield", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, orisa, sigma, dps, reaper, doomfist, support, lucio, moira)},
		{Name: "Double Shield Bastion Bunker", Value: fmt.Sprintf("%s %s %s %s %s %s %s %s %s", tank, orisa, sigma, dps, reaper, bastion, support, baptiste, ana)},
	}
	if _, err := sendMessage(secondPage); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
		return
	}
	return
}

//params currently unused
func removePoll(params []string) {
	channelId := thisSession.ws.events.cachedMessagePayload.ChannelId
	var discordMessageRequest discordMessageRequest
	guildId := thisSession.ws.events.cachedMessagePayload.GuildId
	cachedPoll, ok := pollCache[guildId+channelId]
	if ok {
		delete(pollCache, cachedPoll.Guild+cachedPoll.Channel)
		avatarUrl, _ := thisSession.ws.getUserAvatarOrDefaultUrl(cachedPoll.Creator.Id, cachedPoll.Creator.Avatar, cachedPoll.Creator.Discriminator)
		discordMessageRequest.Embed.Author.Name = cachedPoll.Creator.Username + "#" + cachedPoll.Creator.Discriminator + "´s Poll"
		discordMessageRequest.Embed.Author.IconUrl = avatarUrl
		discordMessageRequest.Embed.Title = "Poll successfully deleted."
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
		discordMessageRequest.Embed.Footer.Text = tipPollCreated
		discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon
		if _, err := sendMessage(discordMessageRequest); err != nil {
			sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
			return
		}
		return
	}

	sendInfoMessageRequest("No Poll was created for this chanel to delete. First create a Poll with !Poll <n>. !Help for more details.")
	return
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
		//Check if everybody responded with ready
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
		discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
		discordMessageRequest.Embed.Footer.Text = infoPollTimeout
		discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon
		discordMessageRequest.Content = "<@" + cachedPollObject.Creator.Id + ">"
		var cachedPollMembers []discordEmbedFieldObject
		for _, val := range cachedPollObject.Members {
			cachedPollMembers = append(cachedPollMembers, discordEmbedFieldObject{Name: val.User.Username, Value: getReadyStatValue(val.Ready, val.Reason), Inline: true})
		}
		discordMessageRequest.Embed.Fields = cachedPollMembers
		if _, err := sendMessage(discordMessageRequest); err != nil {
			sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
			return
		}
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
	guildId := thisSession.ws.events.cachedMessagePayload.GuildId
	channelId := thisSession.ws.events.cachedMessagePayload.ChannelId
	var response string
	if params != nil {
		response = params[0]
	}
	cachedPoll, ok := pollCache[guildId+channelId]
	if ok {
		cachedAuthor := thisSession.ws.events.cachedMessagePayload.Author
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
		discordMessageRequest.Embed.Description = fmt.Sprintf("Poll times out in %.0fmin", cachedPoll.CreationTime.Add(timeout).Sub(time.Now()).Minutes())
		discordMessageRequest.Embed.Title = response + " :expressionless:"
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
		discordMessageRequest.Embed.Footer.Text = tipPollAccept
		discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon
		if _, err := sendMessage(discordMessageRequest); err != nil {
			sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
			return
		}
		checkIfPollIsDone(cachedPoll)
		return
	}
	sendInfoMessageRequest("You have to create a poll first in order to accept or decline. Polls can be created with !Poll <num of members>.")
	return
}

func setUserReady(params []string) {
	var discordMessageRequest discordMessageRequest
	guildId := thisSession.ws.events.cachedMessagePayload.GuildId
	channelId := thisSession.ws.events.cachedMessagePayload.ChannelId

	cachedPoll, ok := pollCache[guildId+channelId]
	if ok {
		cachedAuthor := thisSession.ws.events.cachedMessagePayload.Author
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
		discordMessageRequest.Embed.Description = fmt.Sprintf("Poll times out in %.0fmin", cachedPoll.CreationTime.Add(timeout).Sub(time.Now()).Minutes())
		discordMessageRequest.Embed.Title = ":ok_hand:"
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
		discordMessageRequest.Embed.Footer.Text = tipPollAccept
		discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon
		if _, err := sendMessage(discordMessageRequest); err != nil {
			sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
			return
		}
		checkIfPollIsDone(cachedPoll)
		return
	}
	sendInfoMessageRequest("You have to create a poll first in order to accept or decline. Polls can be created with !Poll <num of members>.")
	return
}

func getReadyStatValue(ready bool, reason string) string {
	if ready {
		return ":white_check_mark: Ready"
	}

	return ":x: Not Ready " + reason
}

func startReadyPoll(params []string) {
	var discordMessageRequest discordMessageRequest
	guildId := thisSession.ws.events.cachedMessagePayload.GuildId
	channelId := thisSession.ws.events.cachedMessagePayload.ChannelId

	cachedPoll, ok := pollCache[guildId+channelId]

	if ok {
		//Existing poll
		//ignore param
		avatarUrl, _ := thisSession.ws.getUserAvatarOrDefaultUrl(cachedPoll.Creator.Id, cachedPoll.Creator.Avatar, cachedPoll.Creator.Discriminator)
		discordMessageRequest.Embed.Author.Name = cachedPoll.Creator.Username + "#" + cachedPoll.Creator.Discriminator + "´s Poll"
		discordMessageRequest.Embed.Author.IconUrl = avatarUrl
		discordMessageRequest.Embed.Title = fmt.Sprintf("Poll times out in %.0fmin", cachedPoll.CreationTime.Add(timeout).Sub(time.Now()).Minutes())
		discordMessageRequest.Embed.Description = fmt.Sprintf("%d out of %d are ready to go!", countReadyMembers(cachedPoll), cachedPoll.Size)
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
		discordMessageRequest.Embed.Footer.Text = tipPollUpdate
		discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon

		var cachedPollMembers []discordEmbedFieldObject
		for _, val := range cachedPoll.Members {
			cachedPollMembers = append(cachedPollMembers, discordEmbedFieldObject{Name: val.User.Username, Value: getReadyStatValue(val.Ready, val.Reason), Inline: true})
		}
		discordMessageRequest.Embed.Fields = cachedPollMembers
		if _, err := sendMessage(discordMessageRequest); err != nil {
			sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
			return
		}
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
		pollCacheObject.Creator = thisSession.ws.events.cachedMessagePayload.Author
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
		discordMessageRequest.Embed.Description = fmt.Sprintf("Accept with !+ or decline with !-. Poll times out in %v", timeout)
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
		discordMessageRequest.Embed.Footer.Text = tipPollCreated
		discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon
		startTimer(timeout, func() {

			cachedPoll, ok := pollCache[guildId+channelId]
			//Check if it is this poll and dont delete a poll which is created after the current poll
			if ok && now.Equal(cachedPoll.CreationTime) {
				delete(pollCache, cachedPoll.Guild+cachedPoll.Channel)
			}
		})
		if _, err := sendMessage(discordMessageRequest); err != nil {
			sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
			return
		}
		return

	} else { //new poll but no param
		sendInfoMessageRequest("You need to specify the size of the poll (e.g. !Poll 5).")
		return
	}
}

func verifyPlatform(val string) bool {

	return utils.ContainsString(platforms, val)
}

func verifyRegion(val string) bool {

	return utils.ContainsString(regions, val)
}

func setGuildConfig(params []string) {
	var discordMessageRequest discordMessageRequest
	if params == nil {
		sendInfoMessageRequest(errorGuildNoParams)
		return
	}

	var guildSettings guildSettingsPersistenceLayer

	var platform string
	var region string
	var prefix string
	for _, param := range params {
		paramStruct := strings.Split(param, "=")

		if len(paramStruct) == 1 {
			sendErrorMessageRequest("Requires value after '=', e.g platform=pc")
			return
		}

		switch paramStruct[0] {
		case "platform":
			if verifyPlatform(paramStruct[1]) {
				platform = paramStruct[1]
			} else {
				sendErrorMessageRequest(errorGuildPlatformNotValid)
				return
			}
		case "region":
			if verifyRegion(paramStruct[1]) {
				region = paramStruct[1]
			} else {
				sendErrorMessageRequest(errorGuildRegionNotValid)
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

	if platform == platformPS || platform == platformXbox {
		region = ""
	}
	if platform == platformPC && region == "" {
		sendErrorMessageRequest(errorGuildReqionRequired)
		return
	}

	//Try load settings
	if err := thisSession.db.getGuildConfig(thisSession.ws.events.cachedMessagePayload.GuildId, &guildSettings); err != nil {
		//Create new
		guildSettings = guildSettingsPersistenceLayer{Platform: platform, Region: region, Prefix: prefix}

		discordMessageRequest.Embed.Author.Name = "Discord Server Config Created"

	} else {
		//Update Existing
		if prefix != "" {
			guildSettings.Prefix = prefix
		}

		if platform != "" {
			guildSettings.Platform = platform
			guildSettings.Region = region
		}

		discordMessageRequest.Embed.Author.Name = "Discord Server Config Updated"
	}

	if err := thisSession.db.setGuildConfig(thisSession.ws.events.cachedMessagePayload.GuildId, &guildSettings); err != nil {
		sendErrorMessageRequest("Error while writing guild config.")
		return
	}

	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
	if _, err := sendMessage(discordMessageRequest); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
		return
	}
}

//noinspection GoUnusedParameter
func getTrainingTimes(params []string) {
	var discordMessageRequest discordMessageRequest
	//Save param as new Training Content in DB
	if params != nil {
		if err := thisSession.db.updateTrainingDates(thisSession.ws.events.cachedMessagePayload.GuildId, trainingDatesPersistenceLayer{params[0]}); err != nil {
			sendErrorMessageRequest(fmt.Sprintf("Error updating Training dates: **%v**\n*%v*\n", params[0], string(err.Error())))
			return
		}
		discordMessageRequest.Embed.Author.Name = "Updated Training days"
		discordMessageRequest.Embed.Description = params[0]
		discordMessageRequest.Embed.Color = 0x970097
		discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
		discordMessageRequest.Embed.Footer.Text = tipMarkup
		discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon
		if _, err := sendMessage(discordMessageRequest); err != nil {
			sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
			return
		}
		return
	}
	var dates trainingDatesPersistenceLayer
	if err := thisSession.db.getTrainingDates(thisSession.ws.events.cachedMessagePayload.GuildId, &dates); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error while retrieving training dates:\n*%v*\n", string(err.Error())))
		return
	}

	discordMessageRequest.Embed.Author.Name = "Training Days"
	discordMessageRequest.Embed.Description = dates.Value
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
	discordMessageRequest.Embed.Footer.Text = tipChangeTraining
	discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon
	if _, err := sendMessage(discordMessageRequest); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
		return
	}
	return
}

//noinspection GoUnusedParameter
func getCurrentlySupportedCommands(params []string) {
	//param unused
	var discordMessageRequest discordMessageRequest
	config := getGuildConfigSave(thisSession.ws.events.cachedMessagePayload.GuildId)
	discordMessageRequest.Embed.Author.Name = "OverwatchTeam Discord Bot - Help"
	discordMessageRequest.Embed.Title = "All currently supported Commands with examples."
	discordMessageRequest.Embed.Description = "If your using Overwatch related commands make sure your profile is set to public"
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
	discordMessageRequest.Embed.Footer.Text = infoUnderConstruction
	discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon
	discordMessageRequest.Embed.Fields = []discordEmbedFieldObject{
		{Name: config.Prefix + "Training", Value: "Displays current Training days"},
		{Name: config.Prefix + "Training <value>", Value: "Updates Training days (e.g. *!Training \"our **new** trainings are ...\"*). Bold, Italic... Style? Check out Discord Markup:arrow_right:" + discordMarkupHelpURL},
		{Name: config.Prefix + "Stats <battletag>", Value: "Displays Player statistics. Player should be registered before *!Register* (e.g. *!Stats Krusher-9911*)"},
		{Name: config.Prefix + "Register <battletag>", Value: "Registers new player. Registered players statistics getting updated automatically every day. (e.g. *!Register Krusher-9911*)"},
		{Name: config.Prefix + "Update <battletag>", Value: "Updates players statistics and stores it or registers the player if not existing. (e.g. *!Update Krusher-9911*)"},
		{Name: config.Prefix + "Config <platform=value region=value> <prefix=value>", Value: "Creates a server config with region and platform to use the Overwatch stats also for Playstation or XboxPlayers. And/or You can also specify a custom prefix for the bot with prefix=value e.g. (*!Config prefix=>*). Supported Platforms are pc, xbl (XBox) or psn (PlayStation)." +
			"Supported Regions are eu,us and asia. Note if your overwatch team is playing on XBox or Playstation, you only need to specify the platform and not the region. (e.g. *!Config platform=psn* for PlayStation or *!Config platform=pc region=us* for PC/US "},
		{Name: config.Prefix + "Poll <number of participants>", Value: fmt.Sprintf("Starts a ready check poll for n players. A poll times out after %d minutes.", timeout/time.Minute)},
		{Name: config.Prefix + "Poll", Value: "Gets the status of the current poll. If everybody is ready a message is created and the creator of the poll gets tagged."},
		{Name: config.Prefix + "+", Value: "Ready."},
		{Name: config.Prefix + "- <reason>", Value: "Not ready. A reason can be passed with the command. (e.g. !- \"need water! Back in 5\"). **Note if your reason is longer then one word you need to put it in \"\"!**"},
		{Name: config.Prefix + "DeletePoll", Value: "Deletes the current Poll."},
		{Name: config.Prefix + "Comps", Value: "Returns all meta comps of Overwatch over the years."},
	}
	if _, err := sendMessage(discordMessageRequest); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error: **%v**\n", string(err.Error())))
		return
	}
	return
}

func getMostPlayedHeroesInMap(carrerStatsMap map[string]heroCareerStats) []kv {
	var gamesPlayedPerHero []kv

	for hero, stats := range carrerStatsMap {
		var gamesPlayed = stats.Game.GamesPlayed
		gamesPlayedPerHero = append(gamesPlayedPerHero, kv{hero, gamesPlayed})
	}

	sort.Slice(gamesPlayedPerHero, func(i, j int) bool {
		return gamesPlayedPerHero[i].Value > gamesPlayedPerHero[j].Value
	})

	return gamesPlayedPerHero
}

type kv struct {
	Key   string
	Value int
}

func getTrendIcon(v1 float32, v2 float32) string {
	sum := v1 - v2
	if sum > 0 {
		return ":chart_with_upwards_trend:"
	} else if sum < 0 {
		return ":chart_with_downwards_trend:"
	} else {
		return ""
	}
}
func getTrendIconI(v1 int, v2 int) string {
	sum := v1 - v2
	if sum > 0 {
		return ":chart_with_upwards_trend:"
	} else if sum < 0 {
		return ":chart_with_downwards_trend:"
	} else {
		return ""
	}
}

func getOverwatchPlayerStats(params []string) {
	//Currently one param supported
	if len(params) != 1 {
		sendErrorMessageRequest("No Battletag specified! Call help for examples and further information.")
		return
	}

	var messageObject discordMessageRequest
	messageObject.Embed.Color = 0x970097
	messageObject.Embed.Author.Name = "Loading player stats now..."
	messageObject.Embed.Description = "Warning: This process can take up to 10 seconds."
	messageObject.Embed.Thumbnail.Url = overwatchIcon

	param := strings.Replace(params[0], "#", "-", 1)
	msg, err := sendMessage(messageObject)

	if err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
		return
	}

	config := getGuildConfigSave(thisSession.ws.events.cachedMessagePayload.GuildId)

	owPlayerLiveStats, err := getPlayerStats(param, config.Platform, config.Region)
	if err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
		return
	}

	if owPlayerLiveStats.Name == "" {
		sendErrorMessageRequest(fmt.Sprintf("Player %s not found for platform %s", param, config.Platform))
		return
	}

	var carrerStatsLive map[string]heroCareerStats
	var topHeroesLive map[string]topHeroStats
	_ = json.Unmarshal(owPlayerLiveStats.CompetitiveStats.CareerStats, &carrerStatsLive)
	_ = json.Unmarshal(owPlayerLiveStats.CompetitiveStats.TopHeroes, &topHeroesLive)
	herosLiveOrdered := getMostPlayedHeroesInMap(carrerStatsLive)

	var owPlayerPersistenceStats owStatsPersistenceLayer

	messageObject.Embed.Footer.Text = tipUpdateProfile
	messageObject.Embed.Footer.IconUrl = overwatchIcon
	if err = thisSession.db.readPlayer(param, &owPlayerPersistenceStats); err != nil {
		//If player is not registered warn and use live stats as persistence stats to avoid weird compared stats
		owPlayerPersistenceStats.OWPlayer = *owPlayerLiveStats
		messageObject.Embed.Footer.Text = fmt.Sprintf("WARNING: The requested player is not registered therefore the statistics containing the data of the whole current season. If you want your daily updated statistics you need to call `!Register %v` once.", param)
		messageObject.Embed.Footer.IconUrl = errorIcon
	}

	//Warning if about to compare account on different platforms
	if owPlayerPersistenceStats.Platform != "" && owPlayerPersistenceStats.Platform != config.Platform {
		sendInfoMessageRequest(fmt.Sprintf("**WARNING:** Platform in saved stats differs from current platform. That means you are about to compare different Accounts.\nIf the Account you are looking for is on platform *__%v__* you need to change the current platform, with the config command.\nIf the Account you are looking for is on platform *__%v__* you need to call the Update Command.", owPlayerPersistenceStats.Platform, config.Platform))
	}

	var carrerStatsPersistent map[string]heroCareerStats
	_ = json.Unmarshal(owPlayerPersistenceStats.OWPlayer.CompetitiveStats.CareerStats, &carrerStatsPersistent)

	var topHeroesPersistent map[string]topHeroStats
	_ = json.Unmarshal(owPlayerPersistenceStats.OWPlayer.CompetitiveStats.CareerStats, &carrerStatsPersistent)

	var winrateAll int
	var winrateToday int

	if owPlayerLiveStats.CompetitiveStats.Games.Played != 0 {
		winrateAll = int(float32(owPlayerLiveStats.CompetitiveStats.Games.Won) / float32(owPlayerLiveStats.CompetitiveStats.Games.Played) * 100.0)
	}
	if owPlayerLiveStats.CompetitiveStats.Games.Played-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Played != 0 {
		winrateToday = int(float32(owPlayerLiveStats.CompetitiveStats.Games.Won-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Won) /
			float32(owPlayerLiveStats.CompetitiveStats.Games.Played-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Played) * 100.0)
	}

	pages := (len(owPlayerLiveStats.Ratings)) + 1
	currentPage := 1

	messageObject.Embed.Color = 0x970097
	messageObject.Embed.Author.Name = owPlayerLiveStats.Name
	messageObject.Embed.Description = fmt.Sprintf("1/%d", pages)
	messageObject.Embed.Author.IconUrl = owPlayerLiveStats.Icon
	messageObject.Embed.Title = "Competitive Statistics"
	messageObject.Embed.Thumbnail.Url = owPlayerLiveStats.RatingIcon
	fields := []discordEmbedFieldObject{
		{Name: "Rating (Combined)", Value: strconv.Itoa(owPlayerLiveStats.Rating) + " SR", Inline: true},
		{Name: "Trend (Combined)", Value: fmt.Sprintf("%+d SR", owPlayerLiveStats.Rating-owPlayerPersistenceStats.OWPlayer.Rating), Inline: true},
		{Name: "Played (all)", Value: strconv.Itoa(owPlayerLiveStats.CompetitiveStats.Games.Played), Inline: true},
		{Name: "Won (all)", Value: fmt.Sprintf("%d  *Win Percentage: %d%%*", owPlayerLiveStats.CompetitiveStats.Games.Won, winrateAll), Inline: true},
		{Name: "Played (today)", Value: strconv.Itoa(owPlayerLiveStats.CompetitiveStats.Games.Played - owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Played), Inline: true},
		{Name: "Won (today)", Value: fmt.Sprintf("%d  *Win Percentage: %d%%*",
			owPlayerLiveStats.CompetitiveStats.Games.Won-owPlayerPersistenceStats.OWPlayer.CompetitiveStats.Games.Won, winrateToday), Inline: true},
	}
	dynamicFields := addDynamicHeroStatsAsEmbedFields(herosLiveOrdered, carrerStatsLive, topHeroesLive, carrerStatsPersistent, topHeroesPersistent)
	fields = append(fields, dynamicFields...)
	messageObject.Embed.Fields = fields

	if _, err = updateMessage(messageObject, msg.Id); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
		return
	}

	currentPage++

	//RoleQ Stats
	if len(owPlayerLiveStats.Ratings) > 0 {
		for _, rating := range owPlayerLiveStats.Ratings {
			//get persistence rating for role
			persistendRatings := owPlayerPersistenceStats.OWPlayer.Ratings
			var persistendRating Rating

			for _, e := range persistendRatings {
				if e.Role == rating.Role {
					persistendRating = e
					break
				}
			}
			var roleQSpecificMessage discordMessageRequest
			roleQSpecificMessage.Embed.Color = 0x970097
			roleQSpecificMessage.Embed.Author.Name = owPlayerLiveStats.Name
			roleQSpecificMessage.Embed.Author.IconUrl = rating.RoleIcon
			roleQSpecificMessage.Embed.Title = fmt.Sprintf("%d/%d", currentPage, pages)
			roleQSpecificMessage.Embed.Thumbnail.Url = rating.RankIcon
			roleQSpecificMessage.Embed.Footer.Text = fmt.Sprintf("Statistics for role %s.", strings.Title(rating.Role))
			roleQSpecificMessage.Embed.Footer.IconUrl = overwatchIcon
			roleQSpecificMessage.Embed.Fields = []discordEmbedFieldObject{
				{
					Name:   "Rating",
					Value:  fmt.Sprintf("%d SR", rating.Level),
					Inline: true,
				},
				{
					Name:   "Trend",
					Value:  fmt.Sprintf("%+d SR", rating.Level-persistendRating.Level),
					Inline: true,
				},
			}
			if _, err := sendMessage(roleQSpecificMessage); err != nil {
				sendErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch role-queue stats for player: **%v**\n*%v*\n", param, string(err.Error())))
				return
			}

			currentPage++
		}
	}

	return
}

func addDynamicHeroStatsAsEmbedFields(herosLiveOrdered []kv, carrerStatsLive map[string]heroCareerStats, topHeroesLive map[string]topHeroStats, carrerStatsPersistent map[string]heroCareerStats, topHeroesPersistent map[string]topHeroStats) []discordEmbedFieldObject {
	//Dynamic hero stats
	var fields []discordEmbedFieldObject
	counter := 1
	for i, v := range herosLiveOrdered {
		if v.Key == "allHeroes" {
			continue
		}

		if i > 6 {
			break
		}

		//Live
		heroStatsLive := carrerStatsLive[v.Key]
		topHeroStatsLive := topHeroesLive[v.Key]

		//Persistent
		heroStatsPersistent := carrerStatsPersistent[v.Key]
		topHeroStatsPersistent := topHeroesPersistent[v.Key]

		//If is support or tank for persistent stats
		roleSpecific := "-"
		var heroSpecificPersistent heroSpecific
		var heroSpecificLive heroSpecific
		_ = json.Unmarshal(heroStatsPersistent.HeroSpecific, &heroSpecificPersistent)
		_ = json.Unmarshal(heroStatsPersistent.HeroSpecific, &heroSpecificLive)
		if heroStatsLive.Assists.HealingDone != 0 {
			healingPerGamePersistent := float32(heroStatsPersistent.Assists.HealingDone) / float32(heroStatsPersistent.Game.GamesPlayed)
			healingPerGameLive := float32(heroStatsLive.Assists.HealingDone) / float32(heroStatsLive.Game.GamesPlayed)
			roleSpecific = fmt.Sprintf("Healing per Game: **%.2f** %s", healingPerGameLive, getTrendIcon(healingPerGameLive, healingPerGamePersistent))
		} else if heroSpecificLive.DamageBlocked != 0 {
			dmgBlockedPerGamePersistent := float32(heroSpecificPersistent.DamageBlocked) / float32(heroStatsPersistent.Game.GamesPlayed)
			dmgBlockedPerGameLive := float32(heroSpecificLive.DamageBlocked) / float32(heroStatsLive.Game.GamesPlayed)
			roleSpecific = fmt.Sprintf("Blocked Dmg per Game: **%.2f** %s", dmgBlockedPerGameLive, getTrendIcon(dmgBlockedPerGameLive, dmgBlockedPerGamePersistent))
		} else if heroStatsLive.Combat.FinalBlows != 0 {
			finalBlowsPerGamePersistent := float32(heroStatsPersistent.Combat.FinalBlows) / float32(heroStatsPersistent.Game.GamesPlayed)
			finalBlowsPerGameLive := float32(heroStatsLive.Combat.FinalBlows) / float32(heroStatsLive.Game.GamesPlayed)
			roleSpecific = fmt.Sprintf("Final blows per Game: **%.2f** %s", finalBlowsPerGameLive, getTrendIcon(finalBlowsPerGameLive, finalBlowsPerGamePersistent))
		}

		weaponAccuracyLive := "-"

		if topHeroStatsLive.WeaponAccuracy > 0 {
			weaponAccuracyLive = fmt.Sprintf("Weapon Accuracy: **%d%%** %s", topHeroStatsLive.WeaponAccuracy, getTrendIcon(float32(topHeroStatsLive.WeaponAccuracy), float32(topHeroStatsPersistent.WeaponAccuracy)))
		}

		//Game stats
		gamesWonLive := topHeroStatsLive.GamesWon
		gamesWonPersistent := topHeroStatsPersistent.GamesWon
		winPercentageLive := topHeroStatsLive.WinPercentage
		winPercentagePersistent := topHeroStatsPersistent.WinPercentage
		dmgDonePerGameLive := float32(heroStatsLive.Combat.DamageDone) / float32(heroStatsLive.Game.GamesPlayed)
		dmgDonePerGamePersistent := float32(heroStatsPersistent.Combat.DamageDone) / float32(heroStatsPersistent.Game.GamesPlayed)

		//averageLive stats
		fields = append(fields, discordEmbedFieldObject{
			Name: fmt.Sprintf("Top Hero #%d %s", counter, HeroIconMap[strings.ToLower(v.Key)]),
			Value: fmt.Sprintf("Games played (all/today): **%v**/**%v**\nGames won (all/today): **%v**/**%v** %s\nWin Percentage: **%d%%** %s\nKD: **%.2f** %s\nDamage per Game: **%.2f** %s\n%s\n%s",
				heroStatsLive.Game.GamesPlayed, heroStatsLive.Game.GamesPlayed-heroStatsPersistent.Game.GamesPlayed,
				gamesWonLive, gamesWonLive-gamesWonPersistent, getTrendIconI(gamesWonLive, gamesWonPersistent),
				winPercentageLive, getTrendIconI(winPercentageLive, winPercentagePersistent),
				topHeroStatsLive.EliminationsPerLife, getTrendIcon(float32(topHeroStatsLive.EliminationsPerLife), float32(topHeroStatsPersistent.EliminationsPerLife)),
				dmgDonePerGameLive, getTrendIcon(dmgDonePerGameLive, dmgDonePerGamePersistent),
				roleSpecific,
				weaponAccuracyLive),
			Inline: true})
		counter++
	}

	return fields
}

func setNewOverwatchPlayer(params []string) {
	//Currently one param supported
	if len(params) != 1 {
		sendErrorMessageRequest("No Battletag specified! Call help for examples and further information.")
		return
	}

	var discordMessageRequest discordMessageRequest
	param := strings.Replace(params[0], "#", "-", 1)

	config := getGuildConfigSave(thisSession.ws.events.cachedMessagePayload.GuildId)

	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Author.Name = "Loading player stats now..."
	discordMessageRequest.Embed.Description = "Warning: This process can take up to 10 seconds."
	discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon

	msg, err := sendMessage(discordMessageRequest)
	if err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
		return
	}

	owPlayerLiveStats, err := getPlayerStats(param, config.Platform, config.Region)
	if err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
		return
	}

	if owPlayerLiveStats.Name == "" {
		sendErrorMessageRequest(fmt.Sprintf("Player %s not found for platform %s", param, config.Platform))
		return
	}

	owStatsPersistenceLayer := owStatsPersistenceLayer{Battletag: param, OWPlayer: *owPlayerLiveStats, Guild: thisSession.ws.events.cachedMessagePayload.GuildId, Platform: config.Platform, Region: config.Region}
	if err = thisSession.db.writePlayer(owStatsPersistenceLayer); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error retrieving Overwatch stats for player: **%v**\n*%v*\n", param, string(err.Error())))
		return
	}
	discordMessageRequest.Embed.Author.Name = owPlayerLiveStats.Name
	discordMessageRequest.Embed.Author.IconUrl = owPlayerLiveStats.Icon
	discordMessageRequest.Embed.Title = "Player added/refreshed"
	discordMessageRequest.Embed.Description = ""
	discordMessageRequest.Embed.Color = 0x970097
	discordMessageRequest.Embed.Thumbnail.Url = overwatchIcon
	discordMessageRequest.Embed.Footer.Text = "Tip: To track your sr for each training, just type !Update " + owPlayerLiveStats.Name + " before each training. After or during the Training you can see your progress with !Stats " + owPlayerLiveStats.Name
	discordMessageRequest.Embed.Footer.IconUrl = overwatchIcon

	if _, err = updateMessage(discordMessageRequest, msg.Id); err != nil {
		sendErrorMessageRequest(fmt.Sprintf("Error while sending Overwatch stats to discord client: **%v**\n*%v*\n", param, string(err.Error())))
		return
	}
	return
}

func sendMessage(message discordMessageRequest) (respMsg discordMessageResponse, err error) {
	respMsg, err = thisSession.ws.sendMessageToChannel(message, thisSession.ws.events.cachedMessagePayload.ChannelId)
	if err != nil {
		return discordMessageResponse{}, err
	}

	return respMsg, nil
}

func updateMessage(message discordMessageRequest, messageId string) (respMsg discordMessageResponse, err error) {
	err = thisSession.ws.updateMessageInChanel(message, thisSession.ws.events.cachedMessagePayload.ChannelId, messageId)
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
	request.Embed.Thumbnail.Url = errorIcon
	request.Embed.Footer.Text = errorFooter
	request.Embed.Footer.IconUrl = overwatchIcon
	_, _ = sendMessage(request)
}

func sendInfoMessageRequest(message string) {
	var request discordMessageRequest
	request.Embed.Color = 0x970097
	request.Embed.Author.Name = "Info"
	request.Embed.Description = message
	request.Embed.Thumbnail.Url = overwatchIcon
	request.Embed.Footer.Text = infoUnderConstruction
	request.Embed.Footer.IconUrl = overwatchIcon
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
