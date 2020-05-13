package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const layoutUS = "January 2, 2006"

var v time.Time

func main() {
	deadline := flag.String("deadline", "", "The deadline for the countdown timer in RFC3339 format (e.g. 2019-12-25T15:00:00+01:00)")
	flag.Parse()

	if *deadline == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("Starting countdown. Deadline=%v", *deadline)

	var err error
	v, err = time.Parse(time.RFC3339, *deadline)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	timeRemaining := getTimeRemaining(v)
	if timeRemaining.t <= 0 {
		log.Println("Countdown reached!")
		postToMMFinal("Today is the Day!! :tada: :grimacing:")
		os.Exit(0)
	}

	postToMM()
	log.Printf("Days: **%d** *Hours*: **%d** *Minutes*: **%d** *Seconds*: **%d**\n", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)

}

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func postToMMFinal(msg string) {
	attach := MMAttachment{}
	attach.AddField(MMField{Title: "CountDown for the release", Value: msg})

	payload := MMSlashResponse{
		Username:    "CloudCountDown",
		IconUrl:     "https://f1.pngfuel.com/png/558/426/391/clock-timer-hourglass-egg-timer-alarm-clocks-countdown-symbol-stopwatch-png-clip-art.png",
		Attachments: []MMAttachment{attach},
	}
	send(payload)
}

func postToMM() {
	timeRemaining := getTimeRemaining(v)
	attach := MMAttachment{}
	attach.AddField(MMField{Title: "CountDown for the Cloud Release", Value: fmt.Sprintf("**%d** Days / **%d** Hours / **%d** Minutes / **%d** Seconds", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)})
	attach.AddField(MMField{Title: "Release Date", Value: v.Format(layoutUS)})

	payload := MMSlashResponse{
		Username:    "CloudCountDown",
		IconUrl:     "https://f1.pngfuel.com/png/558/426/391/clock-timer-hourglass-egg-timer-alarm-clocks-countdown-symbol-stopwatch-png-clip-art.png",
		Attachments: []MMAttachment{attach},
	}
	log.Println("Sending Message to MM")
	send(payload)

}

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
