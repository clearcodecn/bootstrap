package main

import (
	"flag"
	"os"
	"os/signal"
	"tools/core/config"
	"tools/core/cron"
	"tools/core/db"
	"tools/core/server"
)

var (
	cfg string
)

func init() {
	flag.StringVar(&cfg, "c", "config.yaml", "配置文件地址")
}

func main() {
	flag.Parse()

	// 1. 初始化配置
	config.Parse(cfg)

	// 2. 初始化数据库.
	db.InitDatabase()

	stopChan := make(chan struct{})
	// 3. 启动定时任务.
	go cron.Start(stopChan)

	// 4. 启动服务器
	go server.StartServer(stopChan)

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

	close(stopChan)
}
