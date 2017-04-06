package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-xorm/dbweb/modules/setting"
	"github.com/lunny/log"

	"github.com/go-xorm/dbweb/models"
)

var (
	isDebug *bool   = flag.Bool("debug", false, "enable debug mode")
	port    *int    = flag.Int("port", 8989, "listen port")
	https   *bool   = flag.Bool("https", false, "enable https")
	isHelp  *bool   = flag.Bool("help", false, "show help")
	homeDir *string = flag.String("home", defaultHome, "set the home dir which contain templates,static,langs,certs")
)

var (
	defaultHome string
	Version     = "0.2.0329"
	Tags        string
)

func help() {
	fmt.Println("dbweb version", Version)
	fmt.Println()
	flag.PrintDefaults()
}

func exePath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

func main() {
	flag.Parse()

	setting.StaticRootPath = *homeDir
	if setting.StaticRootPath == "" {
		ePath, err := exePath()
		if err != nil {
			panic(err)
		}
		setting.StaticRootPath = filepath.Dir(ePath)
	}

	log.Info("dbweb version", Version)
	log.Info("home dir is", setting.StaticRootPath)
	if len(Tags) > 0 {
		log.Info("build with", Tags)
	}

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
		t.RunTLS(filepath.Join(setting.StaticRootPath, "cert.pem"), filepath.Join(setting.StaticRootPath, "key.pem"), listen)
	} else {
		t.Run(listen)
	}
}
