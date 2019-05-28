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

	command := strings.TrimPrefix(strings.Split(e.cachedMessagePayload.Content, " ")[0], prefix)
	cmd, ok := commandMap[strings.ToLower(command)]
	if !ok {
		return nil
	}

	fmt.Printf("%+v Opcode: %v\n", event.T, event.Op)

	content := strings.Trim(strings.Replace(e.cachedMessagePayload.Content, prefix+command, "", -1), " ")

	re := regexp.MustCompile(`".*?"`)

	loc := re.FindAllString(content, -1)
	for _, val := range loc {
		newVal := strings.Replace(val, " ", "{{@}}", -1)
		content = strings.Replace(content, val, newVal, -1)
	}
	var params []string
	if content != "" {
		params = strings.Split(content, " ")
	}

	for index, val := range params {

		//Try unquoted
		if unquotedVal, err := strconv.Unquote(val); err == nil {
			val = unquotedVal
		}
		params[index] = strings.Replace(val, "{{@}}", " ", -1)
	}

	//run cmd as own go routine to parallelize processing. Need to include sending message and message queue
	cmd(params)
	return nil
}
