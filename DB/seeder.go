package main

import (
	"base/App/Handlers/GORM"
	"base/DB/seeds"
)

func main() {
	dbConnection := GORM.OpenConnection()

	//Seeders =======
	_ = seeds.InitUser(dbConnection)
}
