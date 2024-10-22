package tasks

import (
	"sync"
	"time"
)

func taskSchedule() {
	// taskSchedule_1()
}

func taskSchedule_1() {
	var wgTask sync.WaitGroup
	for {
		wgTask.Add(1) //为计数器设置值

		go func() {

			if timeNow := time.Now().Hour(); timeNow < 6 && timeNow > 0 {

				time.Sleep(30 * time.Minute)
			} else {

				time.Sleep(10 * time.Minute)
			}

			wgTask.Done()
		}()

		wgTask.Wait() //阻塞到计数器的值为0
		time.Sleep(time.Minute)
	}
}
