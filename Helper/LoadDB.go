package Helper

import "os"

var DBConnection, DBName, DBHost, DBPort, DBUsername, DBPassword, DBSSLMode string

func LoadDB() {
	LoadEnv()
	DBConnection = os.Getenv("DB_Connection")
	DBName = os.Getenv("DB_Name")
	DBHost = os.Getenv("DB_Host")
	DBPort = os.Getenv("DB_Port")
	DBUsername = os.Getenv("DB_Username")
	DBPassword = os.Getenv("DB_Password")
	DBSSLMode = os.Getenv("DB_SSLMode")
}
