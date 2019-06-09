package main

import (
	"fmt"
	"time"
)

func startAlarmClock(hour int, minute int, second int, pollingFunc func() error, errorChan chan bool) {
	now := time.Now()
	year, month, day := time.Now().Date()
	alarmTime := time.Date(year, month, day, hour, minute, second, 0, now.Location())

	nowIsAfterAlarm := now.After(alarmTime)

	if nowIsAfterAlarm {
		alarmTime = alarmTime.Add(24 * time.Hour)
	}
	expiresIn := time.Until(alarmTime)

	timer := time.NewTimer(expiresIn)
	fmt.Printf("Alarm: in: %v\n", expiresIn)
	go func() {
		for {
			select {
			case <-timer.C:
				if err := pollingFunc(); err != nil {
					fmt.Printf("Error while running timer function. Err: %v\n", err)
				}
				timer.Reset(24 * time.Hour)
				fmt.Printf("Alarm: reset, expires again in 24h\n")

			case <-errorChan:
				timer.Stop()
				return
			}
		}
	}()
}

func startTimer(wait time.Duration, timerFunc func()) {
	go func() {
		timer := time.NewTimer(wait)
		<-timer.C
		timerFunc()
	}()
}
