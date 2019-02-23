package data

type WeatherJson struct {
	Conditions *[]Conditions `json:"weather"`
	TempData   *TempData     `json:"main"`
	Wind       *Wind         `json:"wind"`
	Clouds     *Clouds       `json:"clouds"`
	Rain       *Rain         `json:"rain"`
	Snow       *Snow         `json:"snow"`
	City       string        `json:"name"`
	Sys        *Sys          `json:"sys"`
}

type Conditions struct {
	Type    string `json:"main"`
	Descrip string `json:"description"`
}

type TempData struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type Wind struct {
	Speed  float64 `json:"speed"`
	Degree int     `json:"deg"`
}

type Clouds struct {
	Percent int `json:"all"`
}

type Rain struct {
	OneHour   float64 `json:"1h"`
	ThreeHour float64 `json:"3h"`
}

type Snow struct {
	OneHour   float64 `json:"1h"`
	ThreeHour float64 `json:"3h"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}
