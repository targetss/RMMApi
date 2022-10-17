package main

/*
func (db *DBObject) PostAircrafts(c *gin.Context) {
	air := Aircrafts{}
	if err := c.BindJSON(air); err != nil {
		log.Println("Error Post to Aircrafts, error: ", err)
	}
	strConn := "INSERT INTO bookings.aircrafts (aircraft_code, model, range) VALUES ($1, $2, &3)"
	result, err := db.DB.Exec(strConn, air.AircraftCode, air.Model, air.Range)
	if err != nil {
		log.Println("Error Insert data to DB, Error: ", err, air.AircraftCode, air.Model, air.Range)
		//return
	}
	count, _ := result.RowsAffected()
	fmt.Println("Добавлено строк: ", count)
}
*/
