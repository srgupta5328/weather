package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/srgupta5328/weather/router"
)

func main() {
	url := "https://api.darksky.net/forecast/05f4fa694bcf8db71cc04082539b23d1/42.3601,-71.0589"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "d975e9db-a842-4834-ab5a-4f73acaf14d1")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Marshal(body)
	fmt.Println(res)

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8300", r))
}
