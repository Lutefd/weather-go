package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type apiConfigData struct {
	OpenWeatherMapKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var config apiConfigData
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return apiConfigData{}, err
	}
	return config, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", weather)

	http.ListenAndServe(":8080", nil)
}
