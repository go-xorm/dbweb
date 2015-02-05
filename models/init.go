package models

import (
	"errors"

	"github.com/lunny/nodb"
	"github.com/lunny/nodb/config"
)

var (
	Db            *nodb.DB
	ErrNotExist   = errors.New("not exist")
	ErrParamError = errors.New("param error")
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

	// add admin
	_, err = GetUserByName("admin")
	if err == ErrNotExist {
		err = AddUser(&User{
			Name:     "admin",
			Password: "admin",
		})
	}

	return err
}
