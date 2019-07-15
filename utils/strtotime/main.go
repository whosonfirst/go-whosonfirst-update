package main

import (
	"flag"
	"log"
	"time"
)

func main() {

	format := flag.String("format", "2006-01-02", "...")
	timezone := flag.String("timezone", "", "...")

	flag.Parse()

	var loc *time.Location

	if *timezone != "" {
	
		l, err := time.LoadLocation(*timezone)

		if err != nil {
			log.Fatal(err)
		}

		loc = l
	}

	for _, dt := range flag.Args() {

		t, err := time.Parse(*format, dt)

		if err != nil {

			log.Fatal(err)
		}

		if loc != nil {
			t = t.In(loc)
		}

		log.Println(dt, t.Unix())
	}

}
