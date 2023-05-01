package fetch

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jbattistella/marketbox/models"
)

func MakeMarketDay(path string) []models.MarketDay {
	csvFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}

	reader := csv.NewReader(csvFile)

	var marketDay []models.MarketDay

	var count int

	for {

		line, error := reader.Read()

		count++

		if count == 1 {
			continue
		}

		if error == io.EOF {
			break

		} else if error != nil {
			log.Fatal(error)
		}

		date, _ := time.Parse("2006/1/2", line[0])
		open, _ := strconv.ParseFloat(line[1], 64)
		last, _ := strconv.ParseFloat(line[2], 64)
		high, _ := strconv.ParseFloat(line[3], 64)
		low, _ := strconv.ParseFloat(line[4], 64)
		rng, _ := strconv.ParseFloat(line[5], 64)
		volume, _ := strconv.ParseFloat(line[6], 64)
		weekPOC, _ := strconv.ParseFloat(line[10], 64)
		monthPOC, _ := strconv.ParseFloat(line[11], 64)
		ytdPOC, _ := strconv.ParseFloat(line[9], 64)
		yearPOC, _ := strconv.ParseFloat(line[7], 64)
		threeyearPOC, _ := strconv.ParseFloat(line[8], 64)

		marketDay = append(marketDay, models.MarketDay{
			Date:   date,
			Open:   open,
			Last:   last,
			High:   high,
			Low:    low,
			Range:  rng,
			Volume: volume,
			POC3yr: threeyearPOC,
			POC1yr: yearPOC,
			POC0yr: ytdPOC,
			POC1wk: weekPOC,
			POC1m:  monthPOC,
			Market: "es",
		})

	}

	return marketDay
}
