package main

import (
	"fmt"
	"time"
)

func main() {

	// makeCSV.MakeintradayCSV("C:/Users/JBattistella/go-workspace/sierradata/es5min.txt")
	// makeCSV.ListHeaders("C:/Users/JBattistella/go-workspace/sierradata/esday.txt")
	// makeCSV.MakeDailyCSV("C:/Users/JBattistella/go-workspace/sierradata/esday.txt")

	// makeCSV.ListHeaders("C:/Users/JBattistella/go-workspace/src/github.com/jbattistella/marketbox/objects/onES1d.csv")

	tm := time.Now().Hour()

	if tm > 4 {
		fmt.Println(tm)
	}

	// query := db.NewDataStore()
	// query.ConnectDB()
	// query.UpdateMarketIntraday()
}
