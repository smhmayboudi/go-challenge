package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/smhmayboudi/go-challenge/model"
)

func main() {
	var waitGroup sync.WaitGroup

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/trade", func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 60*time.Millisecond)
		defer cancel()
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
		instrument := &model.Instrument{}
		// var result []map[string]any
		// db.WithContext(ctx).Raw("ESPECIAL SELECT").Find(&result)
		gormDB.WithContext(ctx).Find(instrument)
		if err := gormDB.Error; err != nil {
			log.Fatalln("error in finding from db:", err)

			return
		}
		res, err := json.Marshal(instrument)
		if err := gormDB.Error; err != nil {
			log.Fatalln("error in marshaling the object:", err)

			return
		}
		w.Write(res)
	})

	httpServer := new(http.Server)
	httpServer.Addr = ":8090"
	httpServer.Handler = mux

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		log.Println("starting the http server")

		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("error in starting http server:", err)

			return
		}
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		<-ctx.Done()

		log.Println("shutting down the http server")

		ctxWT, cancel := context.WithTimeout(
			context.Background(),
			time.Second,
		)
		defer cancel()

		if err := httpServer.Shutdown(ctxWT); err != nil {
			log.Fatalln("error in shutting down http server:", err)

			return
		}
	}()

	waitGroup.Wait()
}
