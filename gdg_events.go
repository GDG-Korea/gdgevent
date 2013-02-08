package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func (e Event) PrintSummary() {
	fmt.Println(e.Start, "~", e.End, e.Title)
}

func FatalIf(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// give strart and end in UTC milliseconds
func getGDGEvents(cid string) []Event {
	base := "https://developers.google.com/events/feed/json"
	requestURL := base + fmt.Sprintf("?group=%s&start=0", cid)
	log.Println(requestURL)

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
	/* cid := "102751345660146384940" // chapter ID of Czech Republic Uber */
	cid := "12714242728066184635" // chapter ID for GDG Golnag Korea
	for _, e := range getGDGEvents(cid) {
		e.PrintSummary()
	}
}