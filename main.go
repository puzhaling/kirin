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

// command syntax: kirin City{,} Country
func main() {
	client := &http.Client {
		Timeout: time.Second*10,
	}

	if len(os.Args[1:]) != 2 {
		log.Fatalf("invalid syntax: command syntax: kirin City{,} Country")
	}
	city, country := parseArguments()

	apikey := "Qk2u387CuUWAKibhWIcqmDm3xDKbaw4t"
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

	weather.Print()
}