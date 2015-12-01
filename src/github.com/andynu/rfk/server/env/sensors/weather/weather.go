package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/env"
	"github.com/andynu/rfk/server/observer"
)

func init() {
	observer.Observe("config.loaded", func(msg interface{}) {
		if config.Config.WeatherUndergroundKey == "" {
			log.Printf("env: sensor: weather [disabled] - No Weather Underground Key in Config")
			return
		}
		var weatherSensor *WeatherSensor
		env.RegisterSensor(weatherSensor)
		log.Printf("env: sensor: weather [enabled] - OK")
	})
}

type WeatherSensor struct {
	LastResponse *WUHourlyResponse
}

func (w *WeatherSensor) Sample() *env.Sample {

	//if w.LastResponse == nil || w.LastResponse.Stale() {
	//	w.LastResponse = pollWeatherUnderground()
	//}
	sample := env.Sample{Timestamp: 0, SensorName: "Weather", Value: "Rain"}
	return &sample
}

func pollWeatherUnderground() *WUHourlyResponse {
	key := config.Config.WeatherUndergroundKey
	location := config.Config.WeatherUndergroundLocation
	hourlyWeatherUrl := "http://api.wunderground.com/api/" + key + "/hourly/q/" + location + ".json"
	fmt.Println(hourlyWeatherUrl)

	res, err := http.Get(hourlyWeatherUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var hourlyWeather *WUHourlyResponse
	err = json.Unmarshal(body, hourlyWeather)
	if err != nil {
		panic(err)
	}
	return hourlyWeather
}

type WUEnglishMetric struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type WUHourlyResponse struct {
	Response struct {
		Version        string `json:"version"`
		Termsofservice string `json:"termsofService"`
		Features       struct {
			Hourly int `json:"hourly"`
		} `json:"features"`
	} `json:"response"`
	HourlyForecast []struct {
		Fcttime struct {
			Hour                   string `json:"hour"`
			HourPadded             string `json:"hour_padded"`
			Min                    string `json:"min"`
			MinUnpadded            string `json:"min_unpadded"`
			Sec                    string `json:"sec"`
			Year                   string `json:"year"`
			Mon                    string `json:"mon"`
			MonPadded              string `json:"mon_padded"`
			MonAbbrev              string `json:"mon_abbrev"`
			Mday                   string `json:"mday"`
			MdayPadded             string `json:"mday_padded"`
			Yday                   string `json:"yday"`
			Isdst                  string `json:"isdst"`
			Epoch                  string `json:"epoch"`
			Pretty                 string `json:"pretty"`
			Civil                  string `json:"civil"`
			MonthName              string `json:"month_name"`
			MonthNameAbbrev        string `json:"month_name_abbrev"`
			WeekdayName            string `json:"weekday_name"`
			WeekdayNameNight       string `json:"weekday_name_night"`
			WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
			WeekdayNameUnlang      string `json:"weekday_name_unlang"`
			WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
			Ampm                   string `json:"ampm"`
			Tz                     string `json:"tz"`
			Age                    string `json:"age"`
			Utcdate                string `json:"UTCDATE"`
		} `json:"FCTTIME"`
		Temp      WUEnglishMetric `json:"temp"`
		Dewpoint  WUEnglishMetric `json:"dewpoint"`
		Condition string          `json:"condition"`
		Icon      string          `json:"icon"`
		IconURL   string          `json:"icon_url"`
		Fctcode   string          `json:"fctcode"`
		Sky       string          `json:"sky"`
		Wspd      WUEnglishMetric `json:"wspd"`
		Wdir      struct {
			Dir     string `json:"dir"`
			Degrees string `json:"degrees"`
		} `json:"wdir"`
		Wx        string          `json:"wx"`
		Uvi       string          `json:"uvi"`
		Humidity  string          `json:"humidity"`
		Windchill WUEnglishMetric `json:"windchill"`
		Heatindex WUEnglishMetric `json:"heatindex"`
		Feelslike struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"feelslike"`
		Qpf  WUEnglishMetric `json:"qpf"`
		Snow WUEnglishMetric `json:"snow"`
		Pop  string          `json:"pop"`
		Mslp WUEnglishMetric `json:"mslp"`
	} `json:"hourly_forecast"`
}

func (r *WUHourlyResponse) Stale() bool {
	return false // TODO: Check times
}
