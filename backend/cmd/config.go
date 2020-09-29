package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Config(){
	driverName := os.Getenv("DRIVER_NAME")
	user := os.Getenv("USER_NAME")
	pass := os.Getenv("PASS")
	tcp := os.Getenv("TCP")
	dbName := os.Getenv("DB_NAME")
	ftm.Println(user+":"+pass+"@"+tcp+"/"+dbName)
}