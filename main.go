package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "demo"
)

type Aircraft struct {
	AircraftCode string `json:"aircraft_code"`
	Model        string `json:"model"`
	Range        int32  `json:"range"`
}

func getAircraft(c *gin.Context) {
	strconn := fmt.Sprintf("user=%v host=%v port=%v password=%v dbname=%v sslmode=disable", user, host, port, password, dbname)
	db, err := sql.Open("postgres", strconn)
	if err != nil {
		log.Print(err)
	}

	defer db.Close()

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

	r := gin.Default()
	r.GET("/aircraft", getAircraft)

	r.Run("localhost:8000")
}
