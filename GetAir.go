package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
)

// Список самолетов
type Aircrafts struct {
	AircraftCode string `json:"aircraft_code"` //Код самолета
	Model        string `json:"model"`         //Модель
	Range        int32  `json:"range"`         //Макс. дальность, км
}

type Routes struct {
	Flight               string   `json:"flight_no"`              //Номер рейса
	DepartureAirport     string   `json:"departure_airport"`      //Код аэропорта отправления
	DepartureAirportName string   `json:"departure_airport_name"` //Название аэропорта отправления
	DepartureCity        string   `json:"departure_city"`         //Город отправления
	ArrivalAirport       string   `json:"arrival_airport"`        //Код аэропорта прибытия
	ArrivalAirportName   string   `json:"arrival_airport_name"`   //Название аэропорта прибытия
	ArrivalCity          string   `json:"arrival_city"`           //Город прибытия
	AircraftCode         string   `json:"aircraft_code"`          //Код самолета, IATA
	Duration             string   `json:"duration"`               //Продолжительность полета
	DaysOfWeek           []string `json:"days_of_week"`           //Дни недели, когда выполняются рейсы

}

func (db *DBObject) GetAircraft(c *gin.Context) {
	airCraft := make([]Aircrafts, 0)
	strConn := "select * FROM bookings.aircrafts"
	rows, err := db.DB.Query(strConn)
	if err != nil {
		log.Println(err)
	}
	rows.Columns()
	for rows.Next() {
		AC := Aircrafts{}
		rows.Scan(&AC.AircraftCode, &AC.Model, &AC.Range)
		airCraft = append(airCraft, AC)
	}

	c.IndentedJSON(200, airCraft)

}

func (db *DBObject) GetRoutes(c *gin.Context) {
	routes := make([]Routes, 0)

	strReq := "select * from bookings.routes"
	res, err := db.DB.Query(strReq)
	if err != nil {
		log.Println(err)
	}

	for res.Next() {
		rt := Routes{}
		res.Scan(&rt.Flight, &rt.DepartureAirport, &rt.DepartureAirportName, &rt.DepartureCity, &rt.ArrivalAirport, &rt.ArrivalAirportName, &rt.ArrivalCity, &rt.AircraftCode, &rt.Duration, (*pq.StringArray)(&rt.DaysOfWeek))
		routes = append(routes, rt)

	}
	c.IndentedJSON(200, routes)
}
