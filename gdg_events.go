package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Event struct {
	End              string `json:"end,omitempty"`
	Description      string `json:"description,omitempty"`
	Title            string `json:"title,omitempty"`
	TemporalRelation string `json:"temporalRelation,omitempty"`
	Start            string `json:"start,omitempty"`
	Link             string `json:"link,omitempty"`
	Location         string `json:"location,omitempty"`
	Id               string `json:"id,omitempty"`
}

const (
	TF_GDG_EVENT = "02 Jan 2006 15:04 -0700"
	TF_CALENDAR  = "20060102"
)

func (e Event) PrintSummary() {
	st, _ := time.Parse(TF_GDG_EVENT, e.Start)
	et, _ := time.Parse(TF_GDG_EVENT, e.End)

	fmt.Println(st.Format(TF_CALENDAR), "~", et.Format(TF_CALENDAR), e.Title)
}

func FatalIf(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getGDGEvents(cid string, start, end time.Time) []Event {
	/* log.Println(start, end) */
	base := "https://developers.google.com/events/feed/json"
	requestURL := base + fmt.Sprintf("?group=%s", cid)
	requestURL += fmt.Sprintf("&start=%d", start.Unix())
	if end.After(start) {
		requestURL += fmt.Sprintf("&end=%d", end.Unix())
	}
	/* log.Println(requestURL) */

	resp, err := http.Get(requestURL)
	FatalIf(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	FatalIf(err)

	var evts []Event
	json.Unmarshal(body, &evts)
	return evts
}

func main() {
	cid := "102751345660146384940" // chapter ID of Czech Republic Uber
	/* cid := "12714242728066184635" // chapter ID for GDG Golnag Korea */

	st, _ := time.Parse(TF_CALENDAR, "20120101")
	et := time.Now()
	for _, e := range getGDGEvents(cid, st.UTC(), et.UTC()) {
		e.PrintSummary()
	}
}
