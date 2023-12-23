package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DataStore struct {
	db *gorm.DB
}

func ConnectDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var url = "postgresql://" + os.Getenv("PGUSER") + ":" + os.Getenv("PGPASS") + "@" + os.Getenv("PGHOST") + ":" + os.Getenv("PGPORT") + "/railway"

	DB, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting:%s\n", err)
		log.Fatal(err)
	}

	log.Println("Sucessfully created the PostgreSQL server!")

	return DB, nil
}

