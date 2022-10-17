package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (db *DBObject) GetAccountsUser(c *gin.Context) {
	AccountUsr := make([]AccountsUser, 0)
	strConn := "select last_login, is_superuser, username, first_name, last_name, email, date_joined, last_login_ip from accounts_user where last_login < now()"
	rows, err := db.DB.Query(strConn)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"status":   "200",
			"response": "Data not found",
		})
	case nil:
		for rows.Next() {
			usr := AccountsUser{}
			rows.Scan(&usr.lastLogin, &usr.isSuperUser, &usr.userName, &usr.firstName, &usr.lastName, &usr.email, &usr.dateJoined, &usr.lastLoginIP)
			AccountUsr = append(AccountUsr, usr)
		}
		c.IndentedJSON(http.StatusOK, AccountUsr)
	default:
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusInternalServerError,
			"response": "Server error",
		})
	}

}

/*
func (db *DBObject) GetRoutes(c *gin.Context) {
	routes := make([]Routes, 0)

	strReq := "select * from bookings.routes"
	res, err := db.DB.Query(strReq)
	if err != nil {
		log.Println(err)
	}

	for res.Next() {
		rt := Routes{}
		res.Scan(&rt.Flight, &rt.DepartureAirport, &rt.DepartureAirportName, &rt.DepartureCity, &rt.ArrivalAirport, &rt.ArrivalAirportName, &rt.ArrivalCity, &rt.AircraftCode, &rt.Duration, (*pq.Int64Array)(&rt.DaysOfWeek)) //ТИП pq... !!!
		routes = append(routes, rt)
	}

	c.IndentedJSON(200, routes)
}
*/
