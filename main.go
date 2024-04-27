package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := http.NewServeMux()

	app.HandleFunc("GET /api/v1/trade", func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 60*time.Millisecond)
		defer cancel()
		dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Tehran"
		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		db, err := gormDB.DB()
		if err != nil {
			panic(err)
		}
		defer func(sqlDB *sql.DB) {
			if err = sqlDB.Close(); err != nil {
				panic(err)
			}
		}(db)
		instrument := &Instrument{}
		gormDB.Find(instrument)
		if err := gormDB.Error; err != nil {
			panic(err)
		}
		res, err := json.Marshal(instrument)
		if err := gormDB.Error; err != nil {
			panic(err)
		}
		w.Write(res)
	})
}
