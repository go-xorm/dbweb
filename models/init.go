package models

import (
	"github.com/lunny/nodb"
	"github.com/lunny/nodb/config"
)

var (
	Db *nodb.DB
)

func Init() error {
	cfg := config.NewConfigDefault()
	cfg.DataDir = "./metas"

	var err error
	// init nosql
	db, err := nodb.Open(cfg)
	if err != nil {
		return err
	}

	// select db
	Db, err = db.Select(0)
	if err != nil {
		return err
	}
	return nil
}
