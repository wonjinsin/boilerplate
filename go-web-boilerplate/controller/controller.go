package controller

import (
	"log"
	"os"
	"pikachu/util"
)

var zlog *util.Logger

func init() {
	zlog, err := util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[controller] err[%s]", err.Error())
		os.Exit(1)
	}
}
