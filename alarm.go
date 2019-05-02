package main

import (
	"time"
	"fmt"
)

func startAlarmClock(hour int, minute int, second int, pollingFunc func() (error)) error {
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
	for {
		<-timer.C
		if err := pollingFunc(); err != nil {
			return err
		}
		timer.Reset(24 * time.Hour)
		fmt.Printf("Alarm: reset, expires again in 24h")
	}

	return nil
}
