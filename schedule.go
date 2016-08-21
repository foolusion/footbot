package footbot

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Week struct {
	XMLName xml.Name `xml:"ss"`
	Week    string   `xml: "w,attr"`
	Year    string   `xml:"y,attr"`
	Type    string   `xml:"t,attr"`
	Games   []Game   `xml:"gms>g"`
}

type Game struct {
	Home     string `xml:"h,attr"`
	Away     string `xml:"v,attr"`
	HomeName string `xml:"hnn,attr"`
	AwayName string `xml:"vnn,attr"`
}

func getSchedule(c *http.Client, year, week, seasonType string) {
	resp, err := c.Get("http://www.nfl.com/ajax/scorestrip?season=" + year + "&seasonType=" + seasonType + "&week=" + week)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	dec := xml.NewDecoder(resp.Body)
	var w Week
	if err := dec.Decode(&w); err != nil {
		log.Fatal(err)
	}

	fmt.Println(w)
}
