package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	//BASE_URL_FORECAST for forcast information
	baseURLForecast = "https://api.darksky.net/forecast/"
)

var (
	lat  = "42.3601"
	long = "71.0589"
)

//HealthCheck ... Health Check Endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Health Check"))
}

//Callback ... Callback url
func Callback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Callback URI"))
}

//GetForecastHandler Gets forecast data for hardcoded latitude and longitude
func GetForecastHandler(w http.ResponseWriter, r *http.Request) {
	darkSkyKey := os.Getenv("DARK_SKY_SECRET")
	url := baseURLForecast + darkSkyKey + "/" + lat + "," + long

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error executing request, %s", err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading body of response")
	}

	json.Marshal(body)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	w.Write(body)

}
