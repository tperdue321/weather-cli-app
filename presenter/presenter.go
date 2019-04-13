package presenter

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tperdue321/weather-cli/data"
)

func DisplayCurrentWeather(resp *data.WeatherJson) {
	printHeaderLineOne()
	displayLineOne(resp)
	fmt.Print("\n\n")
	printHeaderLineTwo()
	displayLineTwo(resp)
	fmt.Print("\n\n")
	printConditionsLineHeader()
	displayConditions(resp.Conditions)
	fmt.Print("\n\n")
	printHeaderLineFour()
	displayLineFour(resp)
	fmt.Print("\n\n")
}

func printHeaderLineOne() {
	headerline1 := []interface{}{
		"City",
		"Country",
		"Sunrise",
		"Sunset",
		"Calculated At",
	}
	line1 := fmt.Sprintf("|%-15s|%-15s|%-15s|%-15s|%-15s|", headerline1...)
	fmt.Println(line1)
}
func printHeaderLineTwo() {
	headerLine2 := []interface{}{
		"Temp",
		"Humidity",
		"Pressure",
	}
	line2 := fmt.Sprintf("|%-15s|%-15s|%-15s|", headerLine2...)
	fmt.Println(line2)
}
func printHeaderLineFour() {
	headerLine4 := []interface{}{
		"Clouds",
		"Wind",
		"Rain 1hr",
		"Rain 3hr",
		"Snow 1hr",
		"Snow 3hr",
	}
	line4 := fmt.Sprintf("|%-15s|%-15s|%-15s|%-15s|%-15s|%-15s|", headerLine4...)
	fmt.Println(line4)
}

func displayLineOne(resp *data.WeatherJson) {
	city := resp.City
	country := resp.Sys.Country
	sunrise := parseTime(resp.Sys.Sunrise)
	sunset := parseTime(resp.Sys.Sunset)
	curTime := parseTime(resp.CurTime)
	line := fmt.Sprintf("|%-15s|%-15s|%-15s|%-15s|%-15s|",
		city, country, sunrise, sunset, curTime)
	fmt.Println(line)
}

func displayLineTwo(resp *data.WeatherJson) {
	tempDataStruct := resp.TempData

	temp := tempDataStruct.Temp
	humidity := tempDataStruct.Humidity
	pressure := tempDataStruct.Pressure

	// these represent deviation
	// tempMin := tempData.TempMin
	// tempMax := tempData.TempMax

	line := fmt.Sprintf("|%-15.2f|%-15d|%-15d|",
		temp, humidity, pressure)
	fmt.Println(line)
}

func printConditionsLineHeader() {
	fmt.Println("Conditions:")
}

func displayConditions(conditionsStructArray []*data.Conditions) {
	length := len(conditionsStructArray)
	// conditions and descriptions arrays of any datatype
	var conditions = make([]interface{}, length)
	var descriptions = make([]interface{}, length)
	for i, condition := range conditionsStructArray {
		conditions[i] = condition.Type
		descriptions[i] = condition.Descrip
	}

	format := "%-30s"
	format = strings.Repeat(format, length)
	conditionsLine := fmt.Sprintf(format, conditions...)
	descriptionLine := fmt.Sprintf(format, descriptions...)
	fmt.Println(conditionsLine)
	fmt.Println(descriptionLine)
}

func displayLineFour(resp *data.WeatherJson) {
	clouds := resp.Clouds.Percent
	windSpeed := resp.Wind.Speed
	windDir := resp.Wind.Degree

	var rain1hr float64
	var rain3hr float64
	if resp.Rain != nil {
		rain1hr = resp.Rain.OneHour
		rain3hr = resp.Rain.ThreeHour
	}

	var snow1hr float64
	var snow3hr float64
	if resp.Snow != nil {
		snow1hr = resp.Snow.OneHour
		snow3hr = resp.Snow.ThreeHour
	}

	padding := 15 - len(strconv.Itoa(clouds))
	speedCharCount := len(strconv.FormatFloat(windSpeed, 'f', 0, 64))
	windDirPadding := 15 - len(strconv.Itoa(windDir)) - speedCharCount - 4

	line := fmt.Sprintf("|%-d%-*s|%-.0f%-s %-d%-*s|%-15.2f|%-15.2f|%-15.2f|%-15.2f|",
		clouds, padding, "%", windSpeed, "mph",
		windDir, windDirPadding, "°",
		rain1hr, rain3hr, snow1hr, snow3hr)
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
