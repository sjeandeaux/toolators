package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type now func() time.Time

var timeNow now

func init() {
	timeNow = time.Now
}

//commandLine the arguments command line
type commandLine struct {
	date string

	stdout io.Writer
	stderr io.Writer
	stdin  *os.File
}

func (c *commandLine) init() {

	//flag
	log.SetPrefix("[d-day]\t")
	log.SetOutput(c.stderr)

	flag.StringVar(&c.date, "date", "2018-05-04", "date yyyy-dd-mm")

	flag.Parse()

}

func (c *commandLine) main() int {
	const hoursByDay = 24
	time1, err := time.Parse(time.RFC3339, c.date+"T00:00:00Z")
	if err != nil {
		fmt.Fprintf(c.stderr, fmt.Sprint(err))
		return 1
	}

	delta := int64(time1.Sub(timeNow()).Hours())
	days := (delta / hoursByDay)
	if delta%hoursByDay == 0 {
		fmt.Fprint(c.stdout, days)
	} else {
		fmt.Fprint(c.stdout, days+1)
	}
	return 0
}
