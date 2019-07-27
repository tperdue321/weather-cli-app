package data

type WeatherJsonV2 struct {
	Currently *Currently `json:"currently"`
	Lat       float64    `json:"latitude"`
	Long      float64    `json:"longitude"`
	TZ        string     `json:"timezone"`
	City      string     `json:"-"`
	State     string     `json:"-"`
}

type Currently struct {
	Time         int64  `json:"time"`
	Summary      string `json:"summary"`
	Icon         string `json:"icon"`
	NearestStorm int64  `json:"nearestStormDistance"`
	//   "precipIntensity": 0.0089,
	//   "precipIntensityError": 0.0046,
	PrecipProb float64 `json:"precipProbability"`
	PrecipType string  `json:"precipType"`
	Temp       float64 `json:"temperature"`
	FeelsLike  float64 `json:"apparentTemperature"`
	DewPoint   float64 `json:"dewPoint"`
	Humidity   float64 `json:"humidity"`
	Pressure   float64 `json:"pressure"`
	WindSpeed  float64 `json:"windSpeed"`
	//   "windGust": 12.03,
	//   "windBearing": 246,
	CloudCover float64 `json:"cloudCover"`
	//   "uvIndex": 1,
	//   "visibility": 9.84,
	//   "ozone": 267.44
}

type Daily struct {
}

type Hourly struct {
}

type Alerts struct {
}

type GeoLoc struct {
	Results []*Results `json:"results"`
}

type Results struct {
	Components *Components `json:"components`
	Formatted  string      `json:"formatted"`
}

type Components struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

// for old api
type WeatherJson struct {
	Conditions []*Conditions `json:"weather"`
	TempData   *TempData     `json:"main"`
	Wind       *Wind         `json:"wind"`
	Clouds     *Clouds       `json:"clouds"`
	Rain       *Rain         `json:"rain"`
	Snow       *Snow         `json:"snow"`
	City       string        `json:"name"`
	Sys        *Sys          `json:"sys"`
	CurTime    int64         `json:"dt"`
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
