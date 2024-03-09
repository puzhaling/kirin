package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
	"log"
)

type Search struct {
	Version int32 				`json: "version"`
	Key string	  				`json: "key"`
	Type string				    `json: "type"`
	Rank int32					`json: "rank"`
	LocalizedName string		`json: "localizedName"`
	Country struct {
		LocalizedName string	`json: "County.LocalizedName"`
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

func main() {
	client := &http.Client {
		Timeout: time.Second*10,
	}

	apikey := "p8C9PIGpwqgUzbwc3KV9iCr68PFINML3"
	q := "Moscow"
	URL := "http://dataservice.accuweather.com/locations/v1/cities/autocomplete?apikey="+apikey+"&q="+q

	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal("redirect function fail / HTTP response fail")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var searches []Search
	err = json.Unmarshal(body, &searches)

	for _, i := range searches {
		i.Print()
	}
}