package WebService

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gaeshi-gaeshi/thermo-smart/HeatersController"
	"github.com/gaeshi-gaeshi/thermo-smart/TemperatureSensorController"
)

// Start is used to setup and start the web service
func Start() {
	http.HandleFunc("/temperature", temperatureHandler)
	http.HandleFunc("/heaters", heatersHandler)
	http.ListenAndServe(":9999", nil)
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		currentTemperature, error := TemperatureSensorController.ReadTemperature()
		if error != nil {
			log.Fatal(error)
		}

		fmt.Fprint(w, currentTemperature)
	} else if r.Method == "POST" {
		// Change target temperature
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func heatersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, HeatersController.GetNumberOfWorkingHeaters())
	} else if r.Method == "POST" {
		bodyBytes, error := ioutil.ReadAll(r.Body)
		if error != nil {
			log.Fatal(error)
		}

		bodyString := string(bodyBytes)
		numberOfHeaters, error := strconv.Atoi(bodyString)
		if error != nil {
			log.Fatal(error)
		}

		HeatersController.SetNumberOfWorkingHeaters(numberOfHeaters)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
