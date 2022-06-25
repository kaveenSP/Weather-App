package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the City : ")
	input, _ := reader.ReadString('\n')

	city := strings.TrimSpace(input)
	apiKEY := "16075fb73e3147dcedf3ebfe6cbac2af"
	gLocation := "https://api.openweathermap.org/geo/1.0/direct?q=" + city + "+&limit=5&appid=" + apiKEY

	res, _ := http.Get(gLocation)

	responseData, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	var jsonData geoLocation

	// valid := json.Valid(responseData)

	json.Unmarshal(responseData, &jsonData)
	// fmt.Printf("%#v \n", jsonData)
	// fmt.Println(jsonData[0])

	weatherURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", jsonData[0].Latitude, jsonData[0].Longitude, apiKEY)

	country := jsonData[0].Country

	// fmt.Println(weatherURL)

	u, _ := url.Parse(weatherURL)

	wRes, _ := http.Get(u.String())

	wResponseData, _ := ioutil.ReadAll(wRes.Body)
	wRes.Body.Close()

	var wJsonData weatherData
	json.Unmarshal(wResponseData, &wJsonData)
	// fmt.Println(wJsonData)
	// final, _ := json.MarshalIndent(wJsonData, "", "\t")
	// fmt.Printf("%s\n", final)

	temp := FarToCel(wJsonData.Main.Temp)
	pressure := wJsonData.Main.Pressure
	humidity := wJsonData.Main.Humidity

	fmt.Println("\nCountry : " + country)
	fmt.Println("\nCity : " + city)
	fmt.Printf("\nTemperature : %f C\n", temp)
	fmt.Printf("\nPressure : %d hpa\n", pressure)
	fmt.Printf("\nHumidity : %d %%\n", humidity)

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
