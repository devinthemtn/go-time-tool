package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	timeAlso := flag.Bool("t", false, "time")
	nextTue := flag.Bool("tue", false, "next Tuesday")
	// must be called after all flags are defined
	flag.Parse()
	// =========
	currentTime := time.Now()
	if *timeAlso {
		fmt.Println(currentTime.Format("2006-Oct-02(15:04:05)"))
	} else if *nextTue {
		nxtTue := nextTuesday(time.Now())
		fmt.Println(nxtTue.Format("2006-Oct-02"))
	} else {
		fmt.Println(currentTime.Format("2006-Oct-02"))
	}
}

func nextTuesday(from time.Time) time.Time {
	daysUntilTuesday := (int(time.Tuesday) - int(from.Weekday()) + 7) % 7
	if daysUntilTuesday == 0 {
		daysUntilTuesday = 7
	}
	return from.AddDate(0, 0, daysUntilTuesday)
}
