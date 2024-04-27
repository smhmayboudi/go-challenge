package main

import "time"

type Instrument struct {
	Id   int
	Name string
}

type Trade struct {
	Id           int
	InstrumentId time.Time
	DateEn       int
	Open         int
	High         int
	Low          int
	Close        int
}
