package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	timeAlso := flag.Bool("t", false, "time")
	nextTue := flag.Bool("tue", false, "next Tuesday")
	compressed := flag.Bool("c", false, "compressed")
	showHelp := flag.Bool("h", false, "help")
	unixtimeflag := flag.Bool("u", false, "unix time stamp")
	// must be called after all flags are defined
	flag.Parse()
	// =========
	if *showHelp {
		fmt.Println("Utillity for showing formated time/date")
		fmt.Println("-h shows this help command")
		fmt.Println("-t shows the current date with the current time")
		fmt.Println("-tue shows the date of the next tuesday from today")
		fmt.Println("-c shows the date time in a compressed format and no special chars")
		fmt.Println("-u shows the unix timestamp")
		os.Exit(0)
	}
	currentTime := time.Now()
	unixtime := currentTime.Unix()
	if *timeAlso {
		fmt.Println(currentTime.Format("2006-Jan-02(15:04:05)"))
	} else if *nextTue {
		nxtTue := nextTuesday(time.Now())
		fmt.Println(nxtTue.Format("2006-Jan-02"))
	} else if *compressed {
		fmt.Println(currentTime.Format("2006Jan02T150405"))
	} else if *unixtimeflag {
		fmt.Println(unixtime)
	} else {
		fmt.Println(currentTime.Format("2006-Jan-02"))
	}
}

func nextTuesday(from time.Time) time.Time {
	daysUntilTuesday := (int(time.Tuesday) - int(from.Weekday()) + 7) % 7
	if daysUntilTuesday == 0 {
		daysUntilTuesday = 7
	}
	return from.AddDate(0, 0, daysUntilTuesday)
}
