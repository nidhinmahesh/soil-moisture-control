package main

import (
	"encoding/json" //for Unmarshal
	"fmt" //for print
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type weather struct {
	Number int `json:"number"`

}

func main() {

	url := "api.openweathermap.org/data/2.5/forecast?id=524901&APPID=7afb6f4a54cc55be8e3adc6fc7a2cb31"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs to avoid long timeouts
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "soil-moisture-control")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	cityname := "name"
	jsonErr := json.Unmarshal(body, &cityname)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(cityname)
}