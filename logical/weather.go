package logical

import "encoding/json"

//UnmarshalWeather wrapper func
func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

//Marshal weather wrapper func
func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Weather struct to hold primary data
type Weather struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timezone  string    `json:"timezone"`
	Currently Currently `json:"currently"`
	Minutely  Minutely  `json:"minutely"`
	Hourly    Hourly    `json:"hourly"`
	Daily     Daily     `json:"daily"`
	Flags     Flags     `json:"flags"`
	Offset    int64     `json:"offset"`
}

//Currently struct holds data about current weather information
type Currently struct {
	Time                 int64       `json:"time"`
	Summary              Summary     `json:"summary"`
	Icon                 Icon        `json:"icon"`
	NearestStormDistance *int64      `json:"nearestStormDistance,omitempty"`
	NearestStormBearing  *int64      `json:"nearestStormBearing,omitempty"`
	PrecipIntensity      float64     `json:"precipIntensity"`
	PrecipProbability    float64     `json:"precipProbability"`
	Temperature          float64     `json:"temperature"`
	ApparentTemperature  float64     `json:"apparentTemperature"`
	DewPoint             float64     `json:"dewPoint"`
	Humidity             float64     `json:"humidity"`
	Pressure             float64     `json:"pressure"`
	WindSpeed            float64     `json:"windSpeed"`
	WindGust             float64     `json:"windGust"`
	WindBearing          int64       `json:"windBearing"`
	CloudCover           float64     `json:"cloudCover"`
	UvIndex              int64       `json:"uvIndex"`
	Visibility           float64     `json:"visibility"`
	Ozone                float64     `json:"ozone"`
	PrecipAccumulation   *float64    `json:"precipAccumulation,omitempty"`
	PrecipType           *PrecipType `json:"precipType,omitempty"`
}

//Daily struct has meta information about daily
type Daily struct {
	Summary string       `json:"summary"`
	Icon    Icon         `json:"icon"`
	Data    []DailyDatum `json:"data"`
}

//DailyDatum ...
type DailyDatum struct {
	Time                        int64   `json:"time"`
	Summary                     string  `json:"summary"`
	Icon                        Icon    `json:"icon"`
	SunriseTime                 int64   `json:"sunriseTime"`
	SunsetTime                  int64   `json:"sunsetTime"`
	MoonPhase                   float64 `json:"moonPhase"`
	PrecipIntensity             float64 `json:"precipIntensity"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime      int64   `json:"precipIntensityMaxTime"`
	PrecipProbability           float64 `json:"precipProbability"`
	TemperatureHigh             float64 `json:"temperatureHigh"`
	TemperatureHighTime         int64   `json:"temperatureHighTime"`
	TemperatureLow              float64 `json:"temperatureLow"`
	TemperatureLowTime          int64   `json:"temperatureLowTime"`
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
	ApparentTemperatureHighTime int64   `json:"apparentTemperatureHighTime"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
	ApparentTemperatureLowTime  int64   `json:"apparentTemperatureLowTime"`
	DewPoint                    float64 `json:"dewPoint"`
	Humidity                    float64 `json:"humidity"`
	Pressure                    float64 `json:"pressure"`
	WindSpeed                   float64 `json:"windSpeed"`
	WindGust                    float64 `json:"windGust"`
	WindGustTime                int64   `json:"windGustTime"`
	WindBearing                 int64   `json:"windBearing"`
	CloudCover                  float64 `json:"cloudCover"`
	UvIndex                     int64   `json:"uvIndex"`
	UvIndexTime                 int64   `json:"uvIndexTime"`
	Visibility                  float64 `json:"visibility"`
	Ozone                       float64 `json:"ozone"`
	TemperatureMin              float64 `json:"temperatureMin"`
	TemperatureMinTime          int64   `json:"temperatureMinTime"`
	TemperatureMax              float64 `json:"temperatureMax"`
	TemperatureMaxTime          int64   `json:"temperatureMaxTime"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMinTime  int64   `json:"apparentTemperatureMinTime"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxTime  int64   `json:"apparentTemperatureMaxTime"`
	PrecipType                  *Icon   `json:"precipType,omitempty"`
}

//Flags strict to hold metadata
type Flags struct {
	Sources        []string `json:"sources"`
	NearestStation float64  `json:"nearest-station"`
	Units          string   `json:"units"`
}

//Hourly struct hold hourly meta data
type Hourly struct {
	Summary string      `json:"summary"`
	Icon    Icon        `json:"icon"`
	Data    []Currently `json:"data"`
}

//Minutely to hold data on minutes
type Minutely struct {
	Summary string          `json:"summary"`
	Icon    Icon            `json:"icon"`
	Data    []MinutelyDatum `json:"data"`
}

//MinutelyDatum is minute meta data
type MinutelyDatum struct {
	Time              int64 `json:"time"`
	PrecipIntensity   int64 `json:"precipIntensity"`
	PrecipProbability int64 `json:"precipProbability"`
}

//Icon is just a representation of the image being served from dark sky api
type Icon string

const (
	//ClearNight icon string
	ClearNight Icon = "clear-night"
	//Cloudy icon string
	Cloudy Icon = "cloudy"
	//Fog icon string
	Fog Icon = "fog"
	//IconRain string
	IconRain Icon = "rain"
	//PartlyCloudyDay icon string
	PartlyCloudyDay Icon = "partly-cloudy-day"
	//PartlyCloudyNight icon string
	PartlyCloudyNight Icon = "partly-cloudy-night"
)

//PrecipType type of percipitation
type PrecipType string

const (
	//Rain ...
	Rain PrecipType = "rain"
	//Snow ...
	Snow PrecipType = "snow"
)

//Summary of weather type
type Summary string

const (
	//Clear ...
	Clear Summary = "Clear"
	//Foggy ...
	Foggy Summary = "Foggy"
	//MostlyCloudy ...
	MostlyCloudy Summary = "Mostly Cloudy"
	//Overcast ...
	Overcast Summary = "Overcast"
	//PartlyCloudy ...
	PartlyCloudy Summary = "Partly Cloudy"
)
