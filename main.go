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

type Headline struct {
	EffectiveDate 		string 			`json:"EffectiveDate"`
	EffectiveEpochDate 	int64    		`json:"EffectiveEpochDate"`
	Severity 			int32           `json:"Severity"`
	Text 				string	 		`json:"Text"`
	Category 			string			`json:"Category"`
	EndDate 			string			`json:"EndDate"`
	EndEpochDate 		int64          	`json:"EndEpochDate"`
	MobileLink 			string          `json:"MobileLink"`
	Link 				string	        `json:"Link"`
}

type Maximum struct {
	Value 		float32        	`json:"Value"`
	Unit 		string         	`json:"Unit"`
	UnitType 	int32      		`json:"UnitType"`
}

type Minimum struct {
	Value 		float32       	`json:"Value"`
	Unit 		string         	`json:"Unit"`
	UnitType 	int32      		`json:"UnitType"`
}

type Temperature struct {
	Minimum Minimum `json:"Minimum"`
	Maximum Maximum `json:"Maximum"`
}

type Day struct {
	Icon 					int32    `json:"Icon"`
	IconPhrase 				string   `json:"IconPhrase"`
	HasPrecipitation 		bool   	 `json:"HasPrecipitation"`
	PrecipitationType 		string	 `json:"PrecipitationType"`
	PrecipitationIntensity  string	 `json:"PrecipitationIntensity"`
}

type Night struct {
	Icon 			 int32         `json:"Icon"`        
	IconPhrase 		 string		   `json:"IconPhrase"`
	HasPrecipitation bool		   `json:"HasPrecipitation"`
}

type DailyForecast struct {
	Date 			string              `json:"Date"`
	EpochDate 		int64             	`json:"EpochDate"`
	Temperature 	Temperature			`json:"Temperature"`
	Day 			Day 				`json:"Day"`
	Night 			Night				`json:"Night"`
	Sources 		[]string			`json:"Sources"`
	MobileLink 		string				`json:"MobileLink"`
	Link 			string				`json:"Link"`
}

type Weather struct {
	Headline 	   Headline 	   `json:"Headline"`
	DailyForecasts []DailyForecast `json:"DailyForecasts"`
}	





type Country struct {
	LocalizedName string	`json:"LocalizedName"`
}

type AdministrativeArea struct {
	ID 			  string	`json:"ID"`
	LocalizedName string	`json:"LocalizedName"`
}

type Search struct {
	Version 				int32 				`json:"Version"`
	Key 					string	  			`json:"Key"`
	Type 					string				`json:"Type"`
	Rank 					int32				`json:"Rank"`
	LocalizedName 			string				`json:"LocalizedName"`
	Country 				Country				`json:"Country"`
	AdministrativeArea 		AdministrativeArea 	`json:"AdministrativeArea"`
}



func (s *Search) Print() {
	fmt.Println("LocalizedName: ", s.LocalizedName)
	fmt.Println("Key: ", s.Key)
	fmt.Println("Country: ", s.Country.LocalizedName)
	fmt.Println()
}

func (w Weather) Print() {
	fmt.Println(w.Headline.EffectiveDate)
	fmt.Println(w.Headline.EffectiveDate)
	fmt.Println(w.Headline.EffectiveDate)
	fmt.Println(w.Headline.Text)
	fmt.Println(w.DailyForecasts[0].Temperature.Minimum.Value, "F")
 	fmt.Println(w.DailyForecasts[0].Temperature.Maximum.Value, "F")
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

	apikey := "Qk2u387CuUWAKibhWIcqmDm3xDKbaw4t"
	URL := "http://dataservice.accuweather.com/locations/v1/cities/autocomplete?apikey="+apikey+"&q="+name

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
			//break
		}
		search.Print()
	}


	URL = "http://dataservice.accuweather.com/forecasts/v1/daily/1day/"+searches[searchIdx].Key+"?apikey="+apikey

	var weather Weather

	resp, err = client.Get(URL)
	if err != nil {
		log.Fatal("redirect function fail / HTTP response fail")
	}

	if resp.StatusCode == http.StatusOK {
		body, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if json.Valid([]byte(body)) {
			fmt.Println("json is valid")

			err = json.Unmarshal([]byte(body), &weather)
			if err != nil {
				log.Fatalf("json encoding fail: %v", err)
			}
		}
	} else {
		log.Fatal("recieved non-200 response status code:" , resp.StatusCode)
	}

	weather.Print()
}