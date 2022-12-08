package GORM

import (
	"gateway_api/Helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func OpenConnection() *gorm.DB {
	Helper.LoadDB()
	dsn := "host=localhost user=" + Helper.DBUsername + " password=" + Helper.DBPassword + " dbname=" + Helper.DBName + " port=" + Helper.DBPort + " sslmode=" + Helper.DBSSLMode + " TimeZone=Asia/Baghdad"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	//sqlDB, err := db.DB()
	//// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//sqlDB.SetMaxIdleConns(100)
	//// SetMaxOpenConns sets the maximum number of open connections to the database.
	//sqlDB.SetMaxOpenConns(5000)
	//// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func CloseConnection(db *gorm.DB) {
	DB, _ := db.DB()
	err := DB.Close()
	if err != nil {
		log.Fatalf("Couldn't close database connection: %s", err)
	}
}
