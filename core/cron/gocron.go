package cron

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func init() {
	gocron.Every(1).Days().At("00:00").Do(helloWorld)
}

func Start(stopChan chan struct{}) {
	ch := gocron.Start()
	<-stopChan
	close(ch)
}

func helloWorld() {
	fmt.Println("hello world")
}
