package footbot

import (
	"net/http"
	"testing"
)

func TestGetSchedule(t *testing.T) {
	client := http.DefaultClient
	tests := []struct {
		year       string
		week       string
		seasonType string
		out        Week
	}{
		{year: "2016", week: "1", seasonType: "PRE"},
		{year: "2016", week: "1", seasonType: "REG"},
	}
	for _, test := range tests {
		getSchedule(client, test.year, test.week, test.seasonType)
	}
}
