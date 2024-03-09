package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
	"log"
	"os"
)

type Weather struct {
	Headline struct {
		EffectiveDate string 				`json: "Headline.EffectiveDate"`
		EffectiveEpochDate int64    		`json: "Headline.EffectiveEpochDate"`
		Severity int32              		`json: "Headline.Severity"`
		Text string	 						`json: "Headline.Text"`
		Category string						`json: "Headline.Category"`
		EndDate string						`json: "Headline.EndDate"`
		EndEpochDate int64          		`json: "Headline.EndEpochDate"`
		MobileLink string           		`json: "Headline.MobileLink"`
		Link string	                		`json: "Headline.Link"`
	}                               		`json: "Headline"`

	DailyForecast struct {
		Date string                 		`json: "DailyForecast.Date"`
		EpochDate int64             		`json: "DailyForecast.EpochDate"`
		Temperature struct {

			Minimum struct {        
				Value float32       		`json: "DailyForecast.Temperature.Minimum.Value"`
				Unit string         		`json: "DailyForecast.Temperature.Minimum.Unit"`
				UnitType int32      		`json: "DailyForecast.Temperature.Minimum.UnitType"`
			}                       		`json: "DailyForecast.Temperature.Minimum"`

			Maximum struct {
				Value float32        		`json: "DailyForecast.Temperature.Maximum.Value"`
				Unit string         		`json: "DailyForecast.Temperature.Maximum.Unit"`
				UnitType int32      		`json: "DailyForecast.Temperature.Maximum.UnitType"`
			}                       		`json: "DailyForecast.Temperature.Maximum"`
		}									`json: "DailyForecast.Temperature"`
		
		Day struct {
			Icon int32              		`json: "DailyForecast.Day.Icon"`
			IconPhrase string       		`json: "DailyForecast.Day.IconPhrase"`
			HasPrecipitation bool   		`json: "DailyForecast.Day.HasPrecipitation"`
			PrecipitationType string		`json: "DailyForecast.Day.PrecipitationType"`
			PrecipitationIntensity string	`json: "DailyForecast.Day.PrecipitationIntensity"`
		}									`json: "DailyForecast.Day"`	

		Night struct {
			Icon int32          			`json: "DailyForecast.Night.Icon"`        
			IconPhrase string				`json: "DailyForecast.Night.IconPhrase"`
			HasPrecipitation bool			`json: "DailyForecast.Night.HasPrecipitation"`
		}									`json: "DailyForecast.Night"`

		Sources []string					`json: "Sources"`
		MobileLink string					`json: "MobileLink"`
		Link string							`json: "Link"`
	}
}	

type Search struct {
	Version int32 				`json: "Version"`
	Key string	  				`json: "Key"`
	Type string				    `json: "Type"`
	Rank int32					`json: "Rank"`
	LocalizedName string		`json: "LocalizedName"`
	Country struct {
		LocalizedName string	`json: "Country.LocalizedName"`
	}							`json: "Country"`
	AdministrativeArea struct {
		ID string				`json: "AdministrativeArea.ID"`
		LocalizedName string	`json: "AdministrativeArea.LocalizedName"`
	}						    `json: "AdministrativeArea"`
}

func (s *Search) Print() {
	fmt.Println("LocalizedName: ", s.LocalizedName)
	fmt.Println("Key: ", s.Key)
	fmt.Println()
}

// command syntax: kirin City{,} Country
func main() {
	client := &http.Client {
		Timeout: time.Second*10,
	}

	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalf("invalid syntax: command syntax: kirin City{,} Country")
	}
	name := args[0]
	country := args[1]

	apikey := "p8C9PIGpwqgUzbwc3KV9iCr68PFINML3"
	q := "Moscow"
	URL := "http://dataservice.accuweather.com/locations/v1/cities/autocomplete?apikey="+apikey+"&q="+q

	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal("redirect function fail / HTTP response fail")
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var searches []Search
	err = json.Unmarshal(body, &searches)


	var searchIdx int
	for idx, search := range searches {
		if search.LocalizedName == name && search.Country.LocalizedName == country {
			searchIdx = idx
			break
		}
	}


	URL = "http://dataservice.accuweather.com/forecasts/v1/daily/1day/"+searches[searchIdx].Key

	var result Weather

	resp, err = client.Get(URL)
	if err != nil {
		log.Fatal("redirect function fail / HTTP response fail")
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	err = json.Unmarshal(body, &result)

	fmt.Println(result.DailyForecast.Temperature.Minimum.Value, "F")
	fmt.Println(result.DailyForecast.Temperature.Maximum.Value, "F")
}