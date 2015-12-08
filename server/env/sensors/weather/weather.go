package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

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
		var weatherSensor WeatherSensor
		env.RegisterSensor(&weatherSensor)
		log.Printf("env: sensor: weather [enabled] - OK")
	})
}

type WeatherSensor struct {
	LastResponse *WUHourlyResponse
}

func (w *WeatherSensor) Sample() []env.Sample {

	now := time.Now()
	forecast := w.updatedWeather(now.Add(time.Hour))

	if forecast.IsEmpty() {
		log.Printf("env: sensor: weather: err: empty forecast")
		return nil
	}

	//log.Printf("env: sensor: weather: sample: %v", forecast.IsEmpty())
	//log.Printf("env: sensor: weather: sample: %v", forecast)

	var samples []env.Sample

	samples = append(samples, env.Sample{
		Timestamp:  now,
		SensorName: "Weather:Condition",
		Value:      forecast.Condition})

	samples = append(samples, env.Sample{
		Timestamp:  now,
		SensorName: "Weather:Temp",
		Value:      forecast.Temp.English})

	samples = append(samples, env.Sample{
		Timestamp:  now,
		SensorName: "Weather:TempFeel",
		Value:      forecast.Feelslike.English})

	return samples
}

func (w *WeatherSensor) updatedWeather(targetTime time.Time) WUHourlyForecast {
	if w.LastResponse == nil || w.LastResponse.stale(targetTime) {
		log.Printf("env: sensor: weather: polling...")
		hourlyWeather, err := w.poll()
		if err != nil {
			log.Printf("env: sensor: weather: poll err: %v", err)
			var nilForecast WUHourlyForecast
			return nilForecast
		}
		w.LastResponse = &hourlyWeather
	}
	return mostRecentForecast(targetTime, *w.LastResponse)
}

func (w *WeatherSensor) poll() (WUHourlyResponse, error) {
	var hourlyWeather WUHourlyResponse
	responseJSON, err := w.fetch()
	if err != nil {
		return hourlyWeather, err
	}
	hourlyWeather, err = w.parse(responseJSON)
	// log.Printf("DEBUG: %v", hourlyWeather)
	if err == nil {
		log.Printf("env: sensor: weather: poll() => success")
	}
	return hourlyWeather, err
}

func (w *WeatherSensor) fetch() ([]byte, error) {
	key := config.Config.WeatherUndergroundKey
	location := config.Config.WeatherUndergroundLocation
	hourlyWeatherUrl := "http://api.wunderground.com/api/" + key + "/hourly/q/" + location + ".json"
	log.Printf("env: sensor: weather: fetch=%v", hourlyWeatherUrl)

	res, err := http.Get(hourlyWeatherUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//log.Printf("DEBUG: body=%v", string(body))
	return body, nil
}

func (w *WeatherSensor) parse(responseJSON []byte) (WUHourlyResponse, error) {
	var hourlyWeather WUHourlyResponse
	err := json.Unmarshal(responseJSON, &hourlyWeather)
	if err != nil {
		return WUHourlyResponse{}, err
	}
	return hourlyWeather, nil
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
	HourlyForecast []WUHourlyForecast `json:"hourly_forecast"`
}

type WUHourlyForecast struct {
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
	Feelslike WUEnglishMetric `json:"feelslike"`
	Qpf       WUEnglishMetric `json:"qpf"`
	Snow      WUEnglishMetric `json:"snow"`
	Pop       string          `json:"pop"`
	Mslp      WUEnglishMetric `json:"mslp"`
}

func (w *WUHourlyForecast) IsEmpty() bool {
	var nilHourlyForecast WUHourlyForecast
	return ((*w) == nilHourlyForecast)
}

func (r *WUHourlyResponse) stale(targetTime time.Time) bool {
	var max time.Time
	for _, forecast := range r.HourlyForecast {
		curTime := epochToTime(forecast.Fcttime.Epoch)
		if curTime.After(max) {
			max = curTime
		}
	}
	lastHour := targetTime.Truncate(time.Hour)
	isStale := (max.After(lastHour))
	log.Printf("env: sensor: weather: cache stale=%t lastHour=%v", isStale, lastHour)

	return isStale
}

func epochToTime(epoch string) time.Time {
	var t time.Time
	eventEpoch, err := strconv.Atoi(epoch)
	if err != nil {
		log.Printf("env: sensor: weather - time parse error %q %v", epoch, err)
		return t
	}
	t = time.Unix(int64(eventEpoch), 0)
	return t
}

// price is right rules (<=)
func mostRecentForecast(targetTime time.Time, response WUHourlyResponse) WUHourlyForecast {
	var max WUHourlyForecast
	var maxTime time.Time

	for _, forecast := range response.HourlyForecast {
		forecastTime := epochToTime(forecast.Fcttime.Epoch)
		if forecastTime.After(targetTime) {
			continue
		}
		if forecastTime.After(maxTime) {
			max = forecast
			maxTime = epochToTime(forecast.Fcttime.Epoch)
		}
	}
	return max
}
