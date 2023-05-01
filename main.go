package main

import (
	"github.com/jbattistella/marketbox/makeCSV"
)

func main() {

	makeCSV.MakeintradayCSV("C:/Users/JBattistella/go-workspace/sierradata/es5min.txt")
	// makeCSV.ListHeaders("C:/Users/JBattistella/go-workspace/sierradata/esday.txt")
	// makeCSV.MakeDailyCSV("C:/Users/JBattistella/go-workspace/sierradata/esday.txt")

	// makeCSV.ListHeaders("C:/Users/JBattistella/go-workspace/sierradata/es5min.txt")

	// tm := time.Now().Hour()

	// if tm > 4 {
	// 	fmt.Println(tm)
	// }

	// query := db.NewDataStore()
	// query.ConnectDB()
	// query.UpdateMarketIntraday()

}
