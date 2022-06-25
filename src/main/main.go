package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/wData", getWeatherData)

	e.Logger.Fatal(e.Start(":8000"))

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("\nEnter the City : ")
	// input, _ := reader.ReadString('\n')
	// city := strings.TrimSpace(input)
}

func getWeatherData(c echo.Context) error {
	city := c.QueryParam("city")
	apiKEY := "16075fb73e3147dcedf3ebfe6cbac2af"
	gLocation := "https://api.openweathermap.org/geo/1.0/direct?q=" + city + "+&limit=5&appid=" + apiKEY

	res, _ := http.Get(gLocation)

	responseData, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var jsonData geoLocation

	json.Unmarshal(responseData, &jsonData)

	weatherURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", jsonData[0].Latitude, jsonData[0].Longitude, apiKEY)

	u, _ := url.Parse(weatherURL)

	wRes, _ := http.Get(u.String())

	wResponseData, _ := ioutil.ReadAll(wRes.Body)
	wRes.Body.Close()

	var wJsonData weatherData
	json.Unmarshal(wResponseData, &wJsonData)

	country := jsonData[0].Country
	temp := fmt.Sprintf("%.2f", FarToCel(wJsonData.Main.Temp))
	pressure := fmt.Sprintf("%d", wJsonData.Main.Pressure)
	humidity := fmt.Sprintf("%d", wJsonData.Main.Humidity)

	return c.JSON(http.StatusOK, map[string]string{
		"country":  country,
		"temp":     temp,
		"pressure": pressure,
		"humidity": humidity,
	})
}

type geoLocation []struct {
	// Name       string   `json:"name"`
	// LocalNames struct{} `json:"local_names,omitempty"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
	// State      struct{} `json:"state"`
}

type weatherData struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
}

func FarToCel(farTemp float64) float64 {
	celTemp := farTemp - 273.15
	return celTemp
}

// valid := json.Valid(responseData)
// fmt.Printf("%#v \n", jsonData)
// fmt.Println(jsonData[0])
// fmt.Println(wJsonData)
// final, _ := json.MarshalIndent(wJsonData, "", "\t")
// fmt.Printf("%s\n", final)
