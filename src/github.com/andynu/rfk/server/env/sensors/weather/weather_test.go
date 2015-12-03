package weather

import (
	"testing"
	"time"
	"github.com/andynu/rfk/server/config"
	"fmt"
)

var JSON string = `
{
  "response": {
  "version":"0.1",
  "termsofService":"http://www.wunderground.com/weather/api/d/terms.html",
  "features": {
  "hourly": 1
  }
	}
		,
	"hourly_forecast": [
		{
		"FCTTIME": {
		"hour": "17","hour_padded": "17","min": "00","min_unpadded": "0","sec": "0","year": "2015","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "30","mday_padded": "30","yday": "333","isdst": "0","epoch": "1448920800","pretty": "5:00 PM EST on November 30, 2015","civil": "5:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
		},
		"temp": {"english": "36", "metric": "2"},
		"dewpoint": {"english": "23", "metric": "-5"},
		"condition": "Mostly Cloudy",
		"icon": "mostlycloudy",
		"icon_url":"http://icons.wxug.com/i/c/k/nt_mostlycloudy.gif",
		"fctcode": "3",
		"sky": "71",
		"wspd": {"english": "4", "metric": "6"},
		"wdir": {"dir": "SSE", "degrees": "158"},
		"wx": "Mostly Cloudy",
		"uvi": "0",
		"humidity": "59",
		"windchill": {"english": "32", "metric": "0"},
		"heatindex": {"english": "-9999", "metric": "-9999"},
		"feelslike": {"english": "32", "metric": "0"},
		"qpf": {"english": "0.0", "metric": "0"},
		"snow": {"english": "0.0", "metric": "0"},
		"pop": "0",
		"mslp": {"english": "30.42", "metric": "1030"}
		}
		,
		{
		"FCTTIME": {
		"hour": "18","hour_padded": "18","min": "00","min_unpadded": "0","sec": "0","year": "2015","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "30","mday_padded": "30","yday": "333","isdst": "0","epoch": "1448924400","pretty": "6:00 PM EST on November 30, 2015","civil": "6:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
		},
		"temp": {"english": "35", "metric": "2"},
		"dewpoint": {"english": "24", "metric": "-4"},
		"condition": "Partly Cloudy",
		"icon": "partlycloudy",
		"icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
		"fctcode": "2",
		"sky": "59",
		"wspd": {"english": "5", "metric": "8"},
		"wdir": {"dir": "S", "degrees": "172"},
		"wx": "Partly Cloudy",
		"uvi": "0",
		"humidity": "62",
		"windchill": {"english": "31", "metric": "-0"},
		"heatindex": {"english": "-9999", "metric": "-9999"},
		"feelslike": {"english": "31", "metric": "-0"},
		"qpf": {"english": "0.0", "metric": "0"},
		"snow": {"english": "0.0", "metric": "0"},
		"pop": "0",
		"mslp": {"english": "30.43", "metric": "1030"}
		}
	]
}
	`

func TestParse(t *testing.T) {
	var ws WeatherSensor
	result, err := ws.parse([]byte(JSON))
	if err != nil {
		t.Errorf("Should be parsable: %v", err)
	}

	betweenTime := time.Unix(1448920801, 0)
	forecast := mostRecentForecast(betweenTime, result)
	if forecast != result.HourlyForecast[0] {
		t.Errorf("Expected different forecast")
	}

	if false == result.stale(betweenTime) {
		t.Errorf("Result should not be stale")
	}

	if true == result.stale(time.Now()) {
		t.Errorf("Result should be stale!")
	}
}

func TestPoll(t *testing.T) {
	if config.Config.WeatherUndergroundKey != "" {
		var ws WeatherSensor
		weather, err := ws.poll()
		if err!= nil {
			t.Errorf("poll err: %v\n", err)
		}
		fmt.Printf("weather: %v\n", weather)
	}
}