package main

import (
	"github.com/go-xorm/dbweb/models"
)

func main() {
	err := models.Init()
	if err != nil {
		panic(err)
	}

	t := InitTango()
	t.Run(":8989")
}
