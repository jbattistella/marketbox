package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jbattistella/marketbox/fetch"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DataStore struct {
	db *gorm.DB
}

func NewDataStore() *DataStore {
	return &DataStore{}
}

func (d *DataStore) ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var url = "postgresql://" + os.Getenv("PGUSER") + os.Getenv("PGPASS") + "@" + os.Getenv("PGHOST") + ":" + os.Getenv("PGPORT") + "/railway"
	var err error
	d.db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting:%s\n", err)
		log.Fatal(err)
	}

	log.Println("Sucessfully created the PostgreSQL server!")
}

func (d *DataStore) UpdateMarketIntraday() {
	marketDayData := fetch.MakeMarketDay("objects/onES1d.csv")

	_ = d.db.Create(&marketDayData)

	log.Println("market data created in database")

}
