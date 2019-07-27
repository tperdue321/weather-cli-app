package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/tperdue321/weather-cli/data"
	Params "github.com/tperdue321/weather-cli/params"
	Presenter "github.com/tperdue321/weather-cli/presenter"
)

const ENV_KEY = "WEATHER_API_KEY"
const ENV_KEY_V2 = "WEATHER_API_KEY_V2"
const GEO_LOC_API_KEY = "GEO_LOC_API_KEY"
const GEO_REVERSE_DECODING_KEY = "GEO_REVERSE_DECODING_KEY"

func main() {
	// var weatherResp data.WeatherJson
	var weatherResp data.WeatherJsonV2
	// location := Flags.SetFlags()
	// apiKey, found := os.LookupEnv(ENV_KEY)
	apiKeyV2, foundV2 := os.LookupEnv(ENV_KEY_V2)

	// if !found {
	// 	panic("Need to set missing ENV variable 'WEATHER_API_KEY' to access the api.")
	// }

	// weatherURL := &url.URL{
	// 	Scheme: "https",
	// 	Host:   "api.openweathermap.org",
	// 	Path:   "/data/2.5/weather",
	// }

	weatherURLV2 := &url.URL{
		Scheme: "https",
		Host:   "api.darksky.net",
	}

	if foundV2 {
		geoApiKey, geoApiKeyEnvSet := os.LookupEnv(GEO_LOC_API_KEY)

		if !geoApiKeyEnvSet {
			panic("Need to set missing ENV variable 'WEATHER_API_KEY' to access the api.")
		}

		latitude, longitude := Params.GeoLookup(geoApiKey)
		weatherURLV2.Path = fmt.Sprintf("/forecast/%s/%s,%s", apiKeyV2, latitude, longitude)
	}
	// weatherURL.RawQuery = Params.SetParams(location, apiKey)
	params := url.Values{}
	params.Set("exclude", "minutely,hourly,flags")
	weatherURLV2.RawQuery = params.Encode()

	respV2, errV2 := http.Get(weatherURLV2.String())

	// resp, err := http.Get(weatherURL.String())
	// if err != nil {
	// 	fmt.Println("weather lookup error:")
	// 	fmt.Println(err)
	// }

	if errV2 != nil {
		fmt.Println("weather lookup error:")
		fmt.Println(errV2)
	}
	defer respV2.Body.Close()
	bodyV2, errV2 := ioutil.ReadAll(respV2.Body)

	if errV2 != nil {
		fmt.Println("error line 70")
		fmt.Println(errV2)
	}

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(bodyV2))
	// fmt.Println("\n\n\n")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	errV2 = json.Unmarshal(bodyV2, &weatherResp)
	if errV2 != nil {
		panic(errV2)
	}
	weatherResp.City, weatherResp.State = getCity(weatherResp.Lat, weatherResp.Long)
	// fmt.Println(weatherResp.Location)
	Presenter.DisplayCurrentWeather(&weatherResp)
}

func getCity(lat, long float64) (string, string) {
	reverseGeoLocURL := &url.URL{
		Scheme: "https",
		Host:   "api.opencagedata.com",
		Path:   "geocode/v1/json",
	}
	if key, found := os.LookupEnv(GEO_REVERSE_DECODING_KEY); found {
		params := url.Values{}
		q := fmt.Sprintf("%v,%v", lat, long)
		params.Set("key", key)
		params.Set("q", q)
		reverseGeoLocURL.RawQuery = params.Encode()

		geoLocResp, err := http.Get(reverseGeoLocURL.String())
		if err != nil {
			panic(err)
		}
		defer geoLocResp.Body.Close()
		body, err := ioutil.ReadAll(geoLocResp.Body)

		var geoLoc data.GeoLoc
		if err := json.Unmarshal(body, &geoLoc); err != nil {
			panic(err)
		}
		comp := geoLoc.Results[0].Components
		return comp.City, comp.State
	}
	return "", ""
}
