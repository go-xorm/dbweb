package main

import (
	"flag"
	"fmt"

	"github.com/go-xorm/dbweb/models"
)

var (
	isDebug *bool = flag.Bool("debug", false, "enable debug mode")
	port    *int  = flag.Int("port", 8989, "listen port")
	https   *bool = flag.Bool("https", false, "enable https")
	isHelp *bool = flag.Bool("help", false, "show help")
)

var (
	version = "0.1"
)

func help() {
	fmt.Println("dbweb version", version)
	fmt.Println()
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if *isHelp {
		help()
		return
	}

	err := models.Init()
	if err != nil {
		panic(err)
	}

	err = InitI18n([]string{"en-US", "zh-CN"})
	if err != nil {
		panic(err)
	}

	t := InitTango(*isDebug)

	listen := fmt.Sprintf(":%d", *port)
	if *https {
		t.RunTLS("./cert.pem", "./key.pem", listen)
	} else {
		t.Run(listen)
	}
}
