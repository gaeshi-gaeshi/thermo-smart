package WebService

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gaeshi-gaeshi/thermo-smart/TemperatureSensorController"
)

// Start is used to setup and start the web service
func Start() {
	http.HandleFunc("/temperature", temperatureHandler)
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
	}
}
