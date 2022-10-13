package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"strings"
)

type Aircraft struct {
	AircraftCode string `json:"aircraft_code"`
	Model        string `json:"model"`
	Range        int32  `json:"range"`
}

func getAircraft(c *gin.Context) {

	air := make([]Aircraft, 0)

	rows, err := db.Query("select * from bookings.aircrafts")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		a := Aircraft{}
		rows.Scan(&a.AircraftCode, &a.Model, &a.Range)
		air = append(air, a)
	}
	for _, val := range air {
		if strings.Contains(val.Model, "Боинг") {
			fmt.Println(val.AircraftCode, val.Model, val.Range)
		}
	}

	c.IndentedJSON(http.StatusOK, air)
}

func main() {
	fmt.Println("DB Project")

	connDB := new(DBObject)

	er := errors.New("")
	connDB.DB, er = CreateConnectDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/aircraft", getAircraft)

	r.Run("localhost:8000")
}
