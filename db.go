package main

import "github.com/nanobox-io/golang-scribble"

type dbSession struct {
	driver *scribble.Driver
}

const (
	Collection = "guilds"
)

func createDB(path string) (*dbSession, error) {
	db, err := scribble.New(path, nil)
	if err != nil {
		return nil, err
	}

	return &dbSession{driver: db}, nil
}

func (d *dbSession) writePlayer(playerStats owStatsPersistenceLayer) error {
	if err := d.driver.Write(Collection, playerStats.Battletag, playerStats.OWPlayer); err != nil {
		return err
	}
	return nil
}

func (d *dbSession) readPlayer(battletag string, playerStats *OWPlayer) error {

	if err := d.driver.Read(Collection, battletag, playerStats); err != nil {
		return err
	}

	return nil
}

type owStatsPersistenceLayer struct {
	Battletag string   `json:"battletag"`
	OWPlayer  OWPlayer `json:"ow_player"`
}
