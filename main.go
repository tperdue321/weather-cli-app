package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/tperdue321/weather-cli/data"
	Flags "github.com/tperdue321/weather-cli/flags"
	Params "github.com/tperdue321/weather-cli/params"
	Presenter "github.com/tperdue321/weather-cli/presenter"
)

const ENV_KEY = "WEATHER_API_KEY"

func main() {
	var weatherResp data.WeatherJson
	location := Flags.SetFlags()
	apiKey, found := os.LookupEnv(ENV_KEY)

	if !found {
		panic("Need to set missing ENV variable 'WEATHER_API_KEY' to access the api.")
	}

	weatherURL := &url.URL{
		Scheme: "https",
		Host:   "api.openweathermap.org",
		Path:   "/data/2.5/weather",
	}
	weatherURL.RawQuery = Params.SetParams(location, apiKey)

	resp, err := http.Get(weatherURL.String())
	if err != nil {
		fmt.Println("weather lookup error:")
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	fmt.Println("\n\n\n")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		fmt.Println(err)
	}
	Presenter.DisplayCurrentWeather(&weatherResp)
}
