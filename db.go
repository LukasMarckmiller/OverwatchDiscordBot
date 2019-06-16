package main

import "github.com/nanobox-io/golang-scribble"

type dbSession struct {
	driver *scribble.Driver
}

const (
	CollectionPlayer   = "player"
	CollectionTraining = "training"
	CollectionGuilds   = "guilds"
)

func createDB(path string) (*dbSession, error) {
	db, err := scribble.New(path, nil)
	if err != nil {
		return nil, err
	}

	return &dbSession{driver: db}, nil
}

func (d *dbSession) writePlayer(playerStats owStatsPersistenceLayer) (err error) {
	if err = d.driver.Write(CollectionPlayer, playerStats.Battletag, playerStats); err != nil {
		return
	}
	return
}

func (d *dbSession) readPlayer(battletag string, playerStats *owStatsPersistenceLayer) (err error) {

	if err = d.driver.Read(CollectionPlayer, battletag, playerStats); err != nil {
		return
	}

	return
}

func (d *dbSession) updateTrainingDates(guild string, content trainingDatesPersistenceLayer) (err error) {
	if err = d.driver.Write(CollectionTraining, guild, content); err != nil {
		return
	}
	return
}

func (d *dbSession) getTrainingDates(guild string, content *trainingDatesPersistenceLayer) (err error) {
	if err = d.driver.Read(CollectionTraining, guild, content); err != nil {
		return
	}
	return
}

func (d *dbSession) setGuildConfig(guild string, content *guildSettingsPersistenceLayer) (err error) {
	if err = d.driver.Write(CollectionGuilds, guild, content); err != nil {
		return
	}
	return
}

func (d *dbSession) getGuildConfig(guild string, content *guildSettingsPersistenceLayer) (err error) {
	if err = d.driver.Read(CollectionGuilds, guild, content); err != nil {
		return
	}
	return
}

type owStatsPersistenceLayer struct {
	Battletag string          `json:"battletag"`
	OWPlayer  owCompleteStats `json:"ow_player"`
	Guild     string          `json:"guild"`
	Platform  string          `json:"platform"`
	Region    string          `json:"region"`
}

type trainingDatesPersistenceLayer struct {
	Value string `json:"value"`
}

type guildSettingsPersistenceLayer struct {
	Region   string `json:"region"`
	Platform string `json:"platform"`
	Prefix   string `json:"prefix"`
}
