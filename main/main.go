package main

import (
	"fmt"
	"thermo-smart/config"
	"thermo-smart/data/models"
	"thermo-smart/data/tiedot"
)

// import (
// 	"fmt"

// 	"github.com/thermo-smart/TemperatureSensorController"
// )

func main() {
	// var temp, error = TemperatureSensorController.ReadTemperature()
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }

	// fmt.Println(temp)

	// dbDirectory := "/tmp/thermo-smart"

	// // (Create if not exist) open a database
	// tiedotDb, err := db.OpenDB(dbDirectory)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // if err := tiedotDb.Create("Temperatures"); err != nil {
	// // 	panic(err)
	// // }

	// temperatures := tiedotDb.Use("Temperatures")

	// temperatures.Insert(map[string]interface{}{"Indication": 36, "Date": "2017-07-07", "Threshold": 37})
	// temperatures.Insert(map[string]interface{}{"Indication": 35, "Date": "2017-07-07", "Threshold": 37})
	// temperatures.Insert(map[string]interface{}{"Indication": 38, "Date": "2017-Jul-07", "Threshold": 37})

	// if err := temperatures.Index([]string{"Indication"}); err != nil {
	// 	panic(err)
	// }

	// if err := temperatures.Index([]string{"Date"}); err != nil {
	// 	panic(err)
	// }

	// var query interface{}
	// json.Unmarshal([]byte(`[{"eq": "36.0", "in": ["Indication"]}]`), &query)

	// query := map[string]interface{}{
	// 	"eq": 36,
	// 	"in": []interface{}{"Indication"},
	// }

	// fmt.Println(query)

	// queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

	// if err := db.EvalQuery(query, temperatures, &queryResult); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(queryResult)
	// // Query results are document IDs
	// for id := range queryResult {
	// 	fmt.Printf("Query returned document ID %d\n", id)
	// }

	// // To use the document itself, simply read it back
	// for id := range queryResult {
	// 	readBack, err := temperatures.Read(id)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Printf("Query returned document %v\n", readBack)

	// 	fmt.Println(readBack["Indication"])
	// }

	// tiedotDb.Close()

	context := tiedot.NewContext(config.App.DbDirectory)
	unit, repository := tiedot.NewUnit(context), tiedot.NewTemperaturesRepository(context)

	unit.Begin()
	// temperature := &models.Temperature{
	// 	Indication: 40,
	// 	Date:       time.Now(),
	// 	Threshold:  43,
	// }

	// fmt.Println(now)

	// repository.Insert(temperature)
	temperatures := repository.Find(repository.InitQuery().Field(models.TemperaturesSchema.Date).Equals("2017-08-13T03:30:20+03:00"))
	unit.Commit()

	fmt.Printf("Temperatures:\n%v", temperatures)
}
