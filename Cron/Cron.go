package Cron

import (
	"fmt"
	"gateway_api/Cron/Tasks"
	"github.com/jasonlvhit/gocron"
)

func Init() {
	go func() {
		err := gocron.Every(1).Seconds().Do(Tasks.DoTask)
		if err != nil {
			fmt.Println(err)
			return
		}
		<-gocron.Start()
	}()
}
