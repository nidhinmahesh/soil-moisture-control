package main

import (
	"encoding/json" 
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type weather struct {
	Main struct {
		Temprature float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
}

func Fetchweb() {

	url := "http://api.openweathermap.org/data/2.5/weather?id=1267360&APPID=7afb6f4a54cc55be8e3adc6fc7a2cb31"

	weatherClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs to avoid long timeouts
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "soil-moisture-control")

	res, getErr := weatherClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	weather0 := weather{}
	jsonErr := json.Unmarshal(body, &weather0)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	temp = weather0.Main.Temprature
	hum = weather0.Main.Humidity	
}