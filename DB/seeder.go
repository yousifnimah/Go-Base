package main

import (
	"gateway_api/App/Handlers/GORM"
	"gateway_api/DB/seeds"
)

func main() {
	dbConnection := GORM.OpenConnection()

	//Seeders =======
	_ = seeds.InitUser(dbConnection)
}
