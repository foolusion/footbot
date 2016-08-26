package footbot

import (
	"context"
	"encoding/xml"
	"net/http"

	"github.com/pkg/errors"
)

type Week struct {
	Week       string `xml: "w,attr"`
	Year       string `xml:"y,attr"`
	SeasonType string `xml:"t,attr"`
	Games      []Game `xml:"gms>g"`
}

type Game struct {
	Home      string `xml:"h,attr"`
	HomeName  string `xml:"hnn,attr"`
	HomeScore int    `xml:"hs,attr"`
	Away      string `xml:"v,attr"`
	AwayName  string `xml:"vnn,attr"`
	AwayScore int    `xml:"vs,attr"`
	Day       string `xml:"d,attr"`
	Time      string `xml:"t,attr"`
}

func getSchedule(ctx context.Context, c *http.Client, year, week, seasonType string) (*Week, error) {
	req, err := http.NewRequest("GET", "http://www.nfl.com/ajax/scorestrip?season="+year+"&seasonType="+seasonType+"&week="+week, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create request")
	}
	req.WithContext(ctx)
	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "getting schedule failed")
	}
	defer resp.Body.Close()
	dec := xml.NewDecoder(resp.Body)
	w := &Week{}
	if err := dec.Decode(w); err != nil {
		return nil, errors.Wrap(err, "decoding xml failed")
	}
	return w, nil
}
