package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
			var wmiInfoStr string

			res.Scan(&pc.ID, &pc.Hostname, &pc.Description, &pc.VersionAgent, &pc.OperatingSystem, &wmiInfoStr, &pc.SiteID)

			var q map[string][][]map[string]interface{}
			if err := json.Unmarshal([]byte(wmiInfoStr), &q); err != nil {
				log.Println(err)
			}
			pc.WMIInfo = make([]map[string]interface{}, 0)
			for key, val := range q {
				var interf map[string]interface{}
				if key == "os" {
					//m := q["os"][0][0]["Version"]

					data := struct {
						VersionSystem  string `json:"Caption"`
						NameDNS        string `json:"CSName"`
						OSArchitecture string `json:"OSArchitecture"`
						//InstallDate    time.Time `json:"InstallDate"`
					}{
						VersionSystem:  fmt.Sprintf("%v", val[0][0]["Caption"]),
						NameDNS:        fmt.Sprintf("%v", val[0][0]["CSName"]),
						OSArchitecture: fmt.Sprintf("%v", val[0][0]["OSArchitecture"]),
						//InstallDate:    fmt.Sprintf("%v", val[0][0]["InstallDate"]),
					}
					unmarsh, err := json.Marshal(&data)
					if err != nil {
						db.log.Write([]byte(err.Error()))
					}
					//var interf map[string]interface{}
					json.Unmarshal(unmarsh, &interf)
					pc.WMIInfo = append(pc.WMIInfo, interf)
					//continue
				}
				if key == "cpu" {
					dataCpu := struct {
						Name        string `json:"Name"`
						Family      string `json:"Family"`
						L2CacheSize string `json:"L2CacheSize"`
						L3CacheSize string `json:"L3CacheSize"`
					}{
						Name:        fmt.Sprintf("%v", val[0][0]["Name"]),
						Family:      fmt.Sprintf("%v", val[0][0]["Family"]),
						L2CacheSize: fmt.Sprintf("%v", val[0][0]["L2CacheSize"]),
						L3CacheSize: fmt.Sprintf("%v", val[0][0]["L3CacheSize"]),
					}
					unmarsh, err := json.Marshal(&dataCpu)
					if err != nil {
						db.log.Write([]byte(err.Error()))
					}
					//var interf map[string]interface{}
					r := json.Unmarshal(unmarsh, &interf)
					if r != nil {
						db.log.Write([]byte(r.Error()))
					}
					pc.WMIInfo = append(pc.WMIInfo, interf)
				}
			}
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
		c.JSON(http.StatusOK, site)
	default:
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusInternalServerError,
			"response": "Server error",
		})
	}
}
