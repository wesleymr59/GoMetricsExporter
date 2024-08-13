package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

// DB is the global database connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	var err error

	dsn := "root:root@tcp(192.168.1.115:3306)/BabyPetControl?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open MySQL connection:", err)
	}

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open connection to MySQL:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	log.Println("Connected to MySQL database successfully!")

}
