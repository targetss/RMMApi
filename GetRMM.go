package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (db *DBObject) GetAccountsUser(c *gin.Context) {
	AccountUsr := make([]AccountsUser, 0)
	strConn := "select last_login, accounts_user.is_superuser, username, first_name, last_name, email, date_joined, last_login_ip, accounts_role.name " +
		"from accounts_user, accounts_role where accounts_user.last_login < now() and accounts_user.role_id = accounts_role.id"
	rows, err := db.DB.Query(strConn)
	if err != nil {
		db.Write([]byte(err.Error()))
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
			rows.Scan(&usr.LastLogin, &usr.IsSuperUser, &usr.UserName, &usr.FirstName, &usr.LastName, &usr.Email, &usr.DateJoined, &usr.LastLoginIP, &usr.RoleName)
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

func (db *DBObject) GetPCToSite(c *gin.Context) {

}

func (db *DBObject) GetListSite(s *gin.Context) {

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
