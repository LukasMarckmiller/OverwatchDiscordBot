package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type events struct {
	cachedMessagePayload discordMessageResponse
}

func (e *events) handleMessageCreate(event discordWebsocketPayloadPresentation) error {
	if err := json.Unmarshal(event.D, &e.cachedMessagePayload); err != nil {
		return err
	}

	//Get default prefix
	prefix := loadPrefixOrDefault(e.cachedMessagePayload.GuildId)
	//Filter non command
	if !strings.HasPrefix(e.cachedMessagePayload.Content, prefix) {
		return nil
	}
	//Filter if requesting event triggered by this bot
	if e.cachedMessagePayload.Author.Id == thisSession.ws.BotUserId {
		return nil
	}
	//Filter if author is a bot to avoid getting triggered by other bots and avoid message flooding
	if e.cachedMessagePayload.Author.Bot {
		return nil
	}

	command := strings.TrimPrefix(strings.Split(e.cachedMessagePayload.Content, " ")[0], prefix)
	cmd, ok := commandMap[strings.ToLower(command)]
	if !ok {
		return nil
	}

	fmt.Printf("%+v Opcode: %v in Guild %s\n", event.T, event.Op, e.cachedMessagePayload.GuildId)

	content := strings.Trim(strings.Replace(e.cachedMessagePayload.Content, prefix+command, "", -1), " ")

	re := regexp.MustCompile(`".*?"`)

	loc := re.FindAllString(content, -1)
	for _, val := range loc {
		newVal := strings.Replace(val, " ", "{{@}}", -1)
		newVal, _ = strconv.Unquote(newVal)
		content = strings.Replace(content, val, newVal, -1)
	}
	var params []string
	if content != "" {
		params = strings.Split(content, " ")
	}

	for index, val := range params {
		params[index] = strings.Replace(val, "{{@}}", " ", -1)
	}

	//run cmd as own go routine to parallelize processing. Need to include sending message and message queue
	cmd(params)
	return nil
}
