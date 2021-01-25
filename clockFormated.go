package utils

import (
	"fmt"
	"strconv"
	"time"
)

// Function for format clock with zero numbers if time is minor than 10 value
func ClockRunnerFormated() {
	hour := strconv.Itoa(time.Now().Hour())
	minute := strconv.Itoa(time.Now().Minute())
	second := strconv.Itoa(time.Now().Second())

	if time.Now().Second() < 10 {
		second = "0" + second
	}

	if time.Now().Minute() < 10 {
		minute = "0" + minute
	}

	if time.Now().Hour() < 10 {
		hour = "0" + hour
	}

	fmt.Printf("\n%v:%v:%v", hour, minute, second)
}
