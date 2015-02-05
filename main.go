package main

import (
	"flag"
	"fmt"

	"github.com/go-xorm/dbweb/models"
)

var (
	isDebug *bool = flag.Bool("debug", false, "enable debug mode")
	port    *int  = flag.Int("port", 8989, "listen port")
)

func main() {
	flag.Parse()

	err := models.Init()
	if err != nil {
		panic(err)
	}

	err = InitI18n([]string{"en-US", "zh-CN"})
	if err != nil {
		panic(err)
	}

	t := InitTango(*isDebug)
	t.Run(fmt.Sprintf(":%d", *port))
}
