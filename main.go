package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tperdue321/weather-cli/data"
	Flags "github.com/tperdue321/weather-cli/flags"
	Params "github.com/tperdue321/weather-cli/params"
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
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		fmt.Println(err)
	}
	displayCurrentWeather(&weatherResp)
}

func displayCurrentWeather(resp *data.WeatherJson) {
	printHeaderLineOne()
	displayLineOne(resp)
	fmt.Print("\n\n")
	printHeaderLineTwo()
	fmt.Print("\n\n")
	printHeaderLineThree()
	fmt.Print("\n\n")
	// fmt.Println(resp)
}

func printHeaderLineOne() {
	headerline1 := []interface{}{
		"City",
		"Country",
		"Sunrise",
		"Sunset",
	}
	line1 := fmt.Sprintf("|%-15s|%-15s|%-15s|%-15s|", headerline1...)
	fmt.Println(line1)
}
func printHeaderLineTwo() {
	headerLine2 := []interface{}{
		"Temp",
		"Temp Min",
		"Temp Max",
		"Humidity",
		"Pressure",
	}
	line2 := fmt.Sprintf("|%-15s|%-15s|%-15s|%-15s|%-15s|", headerLine2...)
	fmt.Println(line2)
}
func printHeaderLineThree() {
	headerLine3 := []interface{}{
		"Conditions",
		"Clouds",
		"Wind",
		"Rain 1hr",
		"Rain 3hr",
		"Snow 1hr",
		"Snow 3hr",
	}
	line3 := fmt.Sprintf("|%-15s|%-15s|%-15s|%-15s|%-15s|%-15s|%-15s|", headerLine3...)
	fmt.Println(line3)
}

func displayLineOne(resp *data.WeatherJson) {
	city := resp.City
	country := resp.Sys.Country
	sunrise := parseTime(resp.Sys.Sunrise)
	sunset := parseTime(resp.Sys.Sunset)
	line := fmt.Sprintf("|%-15s|%-15s|%-15s|%-15s|", city, country, sunrise, sunset)
	fmt.Println(line)
}

func precipitationMmToInches(temp int) {

}

func parseTime(timeStamp int64) string {

	var nanosecs int64 = 0
	utcTimestamp := time.Unix(timeStamp, nanosecs)
	location, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println(err)
	}

	return utcTimestamp.In(location).Format("3:04PM")
}
