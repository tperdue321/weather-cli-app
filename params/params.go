package Params

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	Flags "github.com/tperdue321/weather-cli/flags"
)

const CITY_KEY = "HOME_CITY"
const ZIPCODE_KEY = "HOME_ZIPCODE"
const GEO_LOC_API_KEY = "GEO_LOC_API_KEY"
const GEO_LOOKUP_DOMAIN = "api.ipstack.com"

type geoLoc struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

func SetParams(location *Flags.Location, apiKey string) string {
	params := url.Values{}

	setCity(location, &params)
	setZipcode(location, &params)
	params.Set("APPID", apiKey)
	params.Set("units", "imperial")
	// if params.zip && params.city not set || lat and lon flags present{
	if params.Get("zip") == "" && params.Get("city") == "" ||
		params.Get("lat") != "" && params.Get("lon") != "" {
		setGeoLoc(location, &params)
	}

	return params.Encode()
}

func setCity(location *Flags.Location, params *url.Values) {
	var cityQuery string = ""
	city, cityEnvSet := os.LookupEnv(CITY_KEY)
	if location.City != "" {
		cityQuery = fmt.Sprintf("%s,%s", location.City, location.CountryCode)
	} else if cityEnvSet {
		cityQuery = fmt.Sprintf("%s,%s", city, location.CountryCode)
	}

	if cityQuery != "" {
		params.Set("q", cityQuery)
	}

}

func setZipcode(location *Flags.Location, params *url.Values) {
	var zipcodeQuery string = ""
	zipcode, zipcodeEnvSet := os.LookupEnv(ZIPCODE_KEY)
	if location.Zipcode != "" {
		zipcodeQuery = fmt.Sprintf("%s,%s", location.Zipcode, location.CountryCode)
	} else if zipcodeEnvSet {
		zipcodeQuery = fmt.Sprintf("%s,%s", zipcode, location.CountryCode)
	}

	if zipcodeQuery != "" {
		params.Set("zip", zipcodeQuery)
	}

}

func setGeoLoc(location *Flags.Location, params *url.Values) {

	var latitude string
	var longitude string
	// if an api key exists for geo loc lookup exists we can use it
	geoApiKey, geoApiKeyEnvSet := os.LookupEnv(GEO_LOC_API_KEY)

	// if cmd line args for lat & long are passed in use them
	if location.Latitude != "" && location.Longitude != "" {
		latitude = fmt.Sprintf("%s", location.Latitude)
		longitude = fmt.Sprintf("%s", location.Longitude)
		// else if a geo lock api key is available use that
	} else if geoApiKeyEnvSet {
		latitude, longitude = geoLookup(geoApiKey)
	}

	if latitude != "" && longitude != "" {
		params.Set("lat", latitude)
		params.Set("lon", longitude)
	}

}

func geoLookup(geoApiKey string) (string, string) {

	params := url.Values{"access_key": []string{geoApiKey}}
	geo := new(geoLoc)
	geoURL := &url.URL{
		Scheme:   "http",
		Host:     GEO_LOOKUP_DOMAIN,
		Path:     "/check",
		RawQuery: params.Encode(),
	}

	resp, err := http.Get(geoURL.String())
	if err != nil {
		fmt.Println("geo location error:")
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, geo)
	if err != nil {
		fmt.Println("geo body read err:")
		fmt.Println(err)
	}

	lat := strconv.FormatFloat(geo.Lat, 'f', 2, 64)
	lon := strconv.FormatFloat(geo.Lon, 'f', 2, 64)

	return lat, lon
}
