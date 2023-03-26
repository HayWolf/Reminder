package main

import (
	"github.com/HayWolf/Reminder/store"
	"github.com/HayWolf/Reminder/wechat"
	"github.com/ldigit/config"
	"log"
)

const ConfigPath = "config.yaml"

func main() {

	// 加载配置文件
	cfg, err := config.Load(ConfigPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	config.SetGlobalConfig(cfg)

	// 初始化DB
	store.MigrateTable()

	if err := wechat.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
