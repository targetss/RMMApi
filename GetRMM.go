package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (db *DBObject) GetAccountsUser(c *gin.Context) {
	AccountUsr := make([]AccountsUser, 0)
	strConn := "select last_login, accounts_user.is_superuser, username, first_name, last_name, email, date_joined, last_login_ip, accounts_role.name " +
		"from accounts_user, accounts_role where accounts_user.last_login < now() and accounts_user.role_id = accounts_role.id"
	rows, err := db.DB.Query(strConn)
	if err != nil {
		db.log.Write([]byte(err.Error()))
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

func (db *DBObject) GetInfoComputer(c *gin.Context) {

}

func (db *DBObject) GetPCToSite(c *gin.Context) {
	ids := c.Param("id")
	fmt.Println(ids)
	id, err := strconv.Atoi(ids)
	if err != nil || int(id) < 0 {
		db.log.Write([]byte(err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   http.StatusBadRequest,
			"response": "Incorrect id \"Site\"",
		})
	}
	pcSite := make([]ComputerInfo, 0)
	strconn := "select id, hostname, description, version, operating_system, wmi_detail, site_id  from agents_agent where site_id = $1"

	res, err := db.DB.Query(strconn, id)
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, gin.H{
			"status":   http.StatusNotFound,
			"response": "Content not found",
		})
	case nil:
		for res.Next() {
			var pc ComputerInfo
			res.Scan(&pc.ID, &pc.Hostname, &pc.Description, &pc.VersionAgent, &pc.OperatingSystem, &pc.WMI, &pc.SiteID)
			fmt.Println(pc.ID, pc.WMI)
			pcSite = append(pcSite, pc)
		}
		c.JSON(http.StatusOK, pcSite)
	default:
		db.log.Write([]byte(err.Error()))
	}

}

func (db *DBObject) GetListSite(c *gin.Context) {
	site := make([]Site, 0)

	strRes := "select id, name, client_id, created_by from clients_site"

	res, err := db.DB.Query(strRes)
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"response": "Content not found",
		})
	case nil:
		for res.Next() {
			var ls Site
			res.Scan(&ls.ID, &ls.Name, &ls.ClientID, &ls.CreatedBy)
			site = append(site, ls)
		}
	default:
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusInternalServerError,
			"response": "Server error",
		})
	}
}
