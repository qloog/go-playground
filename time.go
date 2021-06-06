package main

import (
	"fmt"
	"time"
)

func main() {

	var curTime = time.Now()
	endTime := getUserMemberEndTime(curTime, 3)
	fmt.Println(curTime, endTime)
}

func getUserMemberEndTime(startTime time.Time, level int) time.Time {
	endTime := time.Time{}
	// 一个月按31天算
	monthDuration := 24 * 3600 * 31 * time.Second
	switch level {
	// 一个月
	case 1:
		endTime = startTime.Add(monthDuration)
		break
	// 3个月
	case 2:
		endTime = startTime.Add(monthDuration * 3)
		break
	// 半年
	case 3:
		endTime = startTime.Add(monthDuration * 6)
		break
	// 一年
	case 4:
		endTime = startTime.Add(monthDuration * 12)
		break
	// 2年
	case 5:
		endTime = startTime.Add(monthDuration * 24 * 2)
		break
	// 3年
	case 6:
		endTime = startTime.Add(monthDuration * 24 * 3)
		break
	}

	return endTime
}
