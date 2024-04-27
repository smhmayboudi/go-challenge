package model

import "time"

type Instrument struct {
	Id   int
	Name string
}

type Trade struct {
	Id           int
	InstrumentId int
	DateEn       time.Time
	Open         int
	High         int
	Low          int
	Close        int
}
