package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/puzhaling/kirin/backends"
)

func getLocationKey(city, country string, searches []backends.Search) (string, error) {
	for _, s := range searches {
		if s.LocalizedName == city && s.Country.LocalizedName == country {
			return s.Key, nil
		}
	}
	return "", fmt.Errorf("unable to find %q city in %q country", city, country)
}

func parseArguments() (string, string) {
	args := os.Args[1:]

	city, _ := strings.CutSuffix(args[0], ",")
	country := args[1]

	return city, country
}

func main() {
	client := &http.Client {
		Timeout: time.Second*10,
	}

	var city, country string = "", ""

	if len(os.Args[1:]) == 0 {
		for {
			fmt.Print("Weather for city: ")
			n, err := fmt.Scanf("%s %s", &city, &country)

			if err != nil {
				log.Printf("scanning error: %v", err)
			} else if err == nil && n == 2 {
				city, _ = strings.CutSuffix(city, ",")
				break
			} else {
				log.Println("invalid syntax")
			}
		}
	} else {
		city, country = parseArguments()
	}


	apikey := "bqnKx5VQLTPEUTFuCSLbmiz6XLFl9KVO"
	URL := "http://dataservice.accuweather.com/locations/v1/cities/autocomplete?apikey="+apikey+"&q="+city

	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal("redirect function fail / HTTP response fail")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("recieved non-200 response status code:", resp.StatusCode)
	}

	jsonStr, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var searches []backends.Search
	err = json.Unmarshal(jsonStr, &searches)

	locKey, err := getLocationKey(city, country, searches)
	if err != nil {
		log.Fatalf("location key error: %v", err)
	}

	URL = "http://dataservice.accuweather.com/forecasts/v1/daily/1day/"+locKey+"?apikey="+apikey

	var weather backends.Weather

	resp, err = client.Get(URL)
	if err != nil {
		log.Fatal("redirect function fail / HTTP response fail")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		jsonStr, err = ioutil.ReadAll(resp.Body)

		if json.Valid([]byte(jsonStr)) {
			err = json.Unmarshal([]byte(jsonStr), &weather)
			if err != nil {
				log.Fatalf("json encoding fail: %v", err)
			}
		}
	} else {
		log.Fatal("recieved non-200 response status code:", resp.StatusCode)
	}

	weather.Echo()
}