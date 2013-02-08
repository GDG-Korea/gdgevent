package main

import (
	"fmt"
	"github.com/dalinaum/gdgevent/event"
	"os"
	"time"
)

const (
	TF_GDG_EVENT = "02 Jan 2006 15:04 -0700"
	TF_CALENDAR  = "20060102"
)

func main() {
	/* cid := "102751345660146384940" // chapter ID of Czech Republic Uber */
	/* cid := "12714242728066184635" // chapter ID for GDG Golnag Korea */
	if len(os.Args) < 3 {
		fmt.Printf("%s: CHAPTERID YEAR\n", os.Args[0])
		os.Exit(1)
	}

	cid := os.Args[1]
	year := os.Args[2]

	st, _ := time.Parse(TF_CALENDAR, year+"0101")
	et, _ := time.Parse(TF_CALENDAR, year+"1231")
	for _, e := range event.GetGDGEvents(cid, st.UTC(), et.UTC()) {
		e.PrintSummary()
	}
}
