package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/puzhaling/kirin/backends"
)


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

	var searches []backends.Search
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

	var weather backends.Weather

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
		log.Fatal("recieved non-200 response status code:", resp.StatusCode)
	}

	weather.Print()
}