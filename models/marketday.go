package models

import "time"

type MarketDay struct {
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Last   float64
	Range  float64
	Volume float64
	POC3yr float64
	POC1yr float64
	POC0yr float64
	POC1wk float64
	POC1m  float64
}
