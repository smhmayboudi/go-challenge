package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/smhmayboudi/go-challenge/model"
)

func main() {
	wg := sync.WaitGroup{}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	records := flag.Int("records", 10, "number of records to generate, default = 10")

	flag.Parse()

	fmt.Println("records:", *records)

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Tehran"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	db, err := gormDB.DB()
	if err != nil {
		log.Fatalln("error in connection pool:", err)

		return
	}
	defer func(sqlDB *sql.DB) {
		if err = sqlDB.Close(); err != nil {
			log.Fatalln("error in closing the db:", err)

			return
		}
	}(db)

	// Go Routine

	wg.Add(1)

	go Insert(ctx, gormDB, &wg, records)

	wg.Wait()
}

func Insert(ctx context.Context, db *gorm.DB, wg *sync.WaitGroup, records *int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	InstrumentId := rand.Intn(1) + 1
	Open := rand.Int()
	High := rand.Int()
	Low := rand.Int()
	Close := rand.Int()
	trades := make([]model.Trade, *records)
	// Batch The Inputs
	for i := 0; i < *records; i++ {
		trades = append(trades, model.Trade{
			Id:           i,
			InstrumentId: InstrumentId,
			DateEn:       RandomTimestamp(),
			Open:         Open,
			High:         High,
			Low:          Low,
			Close:        Close,
		})
	}
	db.WithContext(ctx).Create(trades)
}

func RandomTimestamp() time.Time {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	return randomNow
}
