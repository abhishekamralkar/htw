package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const weatherTemplate = `Current weather for {{.Name}}:
    Now:         {{.Main.Temp}} {{.Unit}}
    High:        {{.Main.TempMax}} {{.Unit}}
    Low:         {{.Main.TempMin}} {{.Unit}}
`

type data struct {
	Coord    coord    `json:"coord"`
	Main     mainData `json:"main"`
	Timezone int
	Name     string
}

type coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type mainData struct {
	Temp        float64 `json:"temp"`
	FeelsLike   float64 `json:"feels_liks"`
	Min         float64 `json:"temp_min"`
	Max         float64 `json:"temp_max"`
	Pressure    float64 `json:"pressure"`
	Humidity    float64 `json:"humidity"`
	SeaLevel    float64 `json:"sea_level"`
	GroundLevel float64 `json:"grnd_leve"`
}

func getData(city string) float64 {
	var apiKey = os.Getenv("OWM_API_KEY")

	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data_obj := data{}
	jsonErr := json.Unmarshal(body, &data_obj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	lang := data_obj.Coord.Lat

	return lang
}

func main() {
	fmt.Println(getData("pune"))

}
