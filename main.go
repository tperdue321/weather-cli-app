package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	Flags "github.com/tperdue321/weather-cli/flags"
	Params "github.com/tperdue321/weather-cli/params"
)

const ENV_KEY = "WEATHER_API_KEY"

func main() {
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
	fmt.Println(string(body))
}

// func SetParams(location *Flags.Location, apiKey string) string {
// 	params := url.Values{}

// 	if location.City != "" {
// 		cityQuery := fmt.Sprintf("%s,%s", location.City, location.CountryCode)
// 		params.Set("q", cityQuery)
// 	}

// 	if location.Zipcode != "" {
// 		zipcodeQuery := fmt.Sprintf("%s,%s", location.Zipcode, location.CountryCode)
// 		params.Set("zip", zipcodeQuery)
// 	}

// 	params.Set("APPID", apiKey)

// 	return params.Encode()
// }
