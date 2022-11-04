package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type data struct {
	Coord    coord    `json:"coord"`
	Main     mainData `json:"main"`
	Sys      sys      `json:"sys"`
	Timezone int
	Name     string
}

type coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type mainData struct {
	Temp        float64 `json:"temp"`
	FeelsLike   float64 `json:"feels_like"`
	Min         float64 `json:"temp_min"`
	Max         float64 `json:"temp_max"`
	Pressure    float64 `json:"pressure"`
	Humidity    float64 `json:"humidity"`
	SeaLevel    float64 `json:"sea_level"`
	GroundLevel float64 `json:"grnd_leve"`
}

type sys struct {
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

func main() {
	var apiKey = os.Getenv("OWM_API_KEY")

	city := flag.String("city", "Pune", "Enter the city name.")
	units := flag.String("units", "kelvin", "Supports Kelvin(standard), Celcius(metric), Fahrenheit(imperial)")

	flag.Parse()
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + *city + "&appid=" + apiKey + "&units=" + *units

	const weatherTemplate = `Current weather in {{.Name}}, {{.Sys.Country}}:
    Now:         {{.Main.FeelsLike}} 
    High:        {{.Main.Max}}
    Low:         {{.Main.Min}}
    Pressure:    {{.Main.Pressure}}
    Humidity:    {{.Main.Humidity}}

`

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
		return
	}

	//create a new template with some name
	tmpl := template.New("weather")

	//parse some content and generate a template
	tmpl, err = tmpl.Parse(weatherTemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	err1 := tmpl.Execute(os.Stdout, data_obj)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}

}
