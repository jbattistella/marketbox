package makeCSV

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type marketIntraday struct {
	Date           string
	Time           string
	Open           string
	High           string
	Low            string
	Last           string
	Range          string
	Volume         string
	Trades         string
	BidVolume      string
	AskVolume      string
	Delta          string
	RelativeVolume string
	IBhigh         string
	IBmid          string
	IBlow          string
	IBhighExt1     string
	IBhighExt2     string
	IBlowExt1      string
	IBlowExt2      string
	USPOC          string
	USVWAP         string
	ONPOC          string
	GlobexVWAP     string
}
type marketDay struct {
	Date   string
	Open   string
	High   string
	Low    string
	Last   string
	Range  string
	Volume string
	Delta  string
	POC3yr string
	POC1yr string
	POCytd string
	POC1wk string
	POC1m  string
}

func MakeintradayCSV(t string) {
	csvFile, _ := os.Open(t)
	reader := csv.NewReader(csvFile)
	reader.TrimLeadingSpace = true
	var es2 []marketIntraday

	for {
		var counter int

		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		bidVolume, _ := strconv.Atoi(line[11])
		askVolume, _ := strconv.Atoi(line[12])
		delta := askVolume - bidVolume
		d := fmt.Sprintf("%d", delta)

		if counter == 0 {
			d = "Delta"
		}
		es2 = append(es2, marketIntraday{
			Date:           line[0],
			Time:           line[1],
			Open:           line[2],
			Last:           line[5],
			High:           line[3],
			Low:            line[4],
			Volume:         line[6],
			Trades:         line[7],
			BidVolume:      line[11],
			AskVolume:      line[12],
			Delta:          d,
			RelativeVolume: line[17],
			IBhigh:         line[23],
			IBmid:          line[24],
			IBlow:          line[25],
			IBhighExt1:     line[21],
			IBhighExt2:     line[22],
			IBlowExt1:      line[27],
			IBlowExt2:      line[28],
			USPOC:          line[30],
			USVWAP:         line[31],
			ONPOC:          line[32],
			GlobexVWAP:     line[17],
		})
		counter++
	}

	file, err2 := os.Create("objects/ES5min.csv")
	defer file.Close()
	if err2 != nil {
		log.Fatalln("failed to open file", err2)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	for _, d := range es2 {
		row := []string{
			d.Date,
			d.Time,
			d.Last,
			d.High,
			d.Low,
			d.Volume,
			d.Trades,
			d.BidVolume,
			d.AskVolume,
			d.Delta,
			d.RelativeVolume,
			d.IBhigh,
			d.IBmid,
			d.IBlow,
			d.IBhighExt1,
			d.IBhighExt2,
			d.IBlowExt1,
			d.IBlowExt2,
			d.USPOC,
			d.USVWAP,
			d.ONPOC,
			d.GlobexVWAP,
		}

		if err3 := w.Write(row); err3 != nil {
			log.Fatalln("error writing record to file", err3)
		}
	}
}

func MakeDailyCSV(t string) {

	csvFile, _ := os.Open(t)
	reader := csv.NewReader(csvFile)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1
	var es2 []marketDay
	var counter int
	for {

		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		high, _ := strconv.ParseFloat(line[3], 64)
		low, _ := strconv.ParseFloat(line[4], 64)
		onrange := high - low
		rng := fmt.Sprintf("%0.02f", onrange)

		if counter == 0 {
			rng = "Range"
		}
		counter++
		es2 = append(es2, marketDay{
			Date:   line[0],
			Open:   line[2],
			Last:   line[5],
			High:   line[3],
			Low:    line[4],
			Range:  rng,
			Volume: line[6],
			POC1wk: line[29],
			POC1m:  line[25],
			POCytd: line[13],
			POC1yr: line[33],
			POC3yr: line[37],
		})
	}

	file, err2 := os.Create("objects/onES1d.csv")
	defer file.Close()
	if err2 != nil {
		log.Fatalln("failed to open file", err2)
	}
	w := csv.NewWriter(file)

	defer w.Flush()

	for _, d := range es2 {
		row := []string{
			d.Date,
			d.Open,
			d.Last,
			d.High,
			d.Low,
			d.Range,
			d.Volume,
			d.POC1yr,
			d.POC3yr,
			d.POCytd,
			d.POC1wk,
			d.POC1m,
		}

		if err3 := w.Write(row); err3 != nil {
			log.Fatalln("error writing record to file", err3)
		}
	}

}
