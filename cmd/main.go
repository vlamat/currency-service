package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/vlamat/currency-service/internal/db"
	"github.com/vlamat/currency-service/internal/handlers"
	"github.com/vlamat/currency-service/internal/repository"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

func main() {
	// Подлючение к БД
	dbConn, err := db.Connect("root", "rootpassword", "localhost", 3306, "currency_db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	// Создание таблицы
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS currency_rates (
            id INT AUTO_INCREMENT PRIMARY KEY,
            cur_id INT NOT NULL,
            date DATE NOT NULL,
            cur_abbreviation VARCHAR(10) NOT NULL,
            cur_scale INT NOT NULL,
            cur_name VARCHAR(100) NOT NULL,
            cur_official_rate DECIMAL(10,4) NOT NULL
        )
    `
	_, err = dbConn.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Создание маршрутов
	r := mux.NewRouter()

	// Создание обработчика
	h := &handlers.Handler{DB: dbConn}

	// GET http://localhost:8080/rates вернёт все собранные записи.
	r.HandleFunc("/rates", h.GetAllRatesHandler).Methods("GET")

	// GET http://localhost:8080/rates/day?date=2025-01-03 вернёт записи на конкретный день (формат YYYY-MM-DD)
	r.HandleFunc("/rates/day", h.GetRatesByDateHandler).Methods("GET")

	// StartScheduler(dbConn) запускает планировщик, который каждый день в 01:00 собирает курсы валют

	// Старт сервера
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func StartScheduler(db *sql.DB) {
	c := cron.New()
	// Каждый день в 01:00
	c.AddFunc("0 1 * * *", func() {
		rates, err := FetchRatesFromNBRB()
		if err != nil {
			log.Println("Error fetching rates:", err)
			return
		}

		// Преобразуем и записываем в БД
		var toInsert []repository.CurrencyRate
		for _, r := range rates {
			dateParsed, err := time.Parse("2006-01-02T15:04:05", r.Date)
			if err != nil {
				log.Println("Error parsing date:", err)
				continue
			}
			toInsert = append(toInsert, repository.CurrencyRate{
				CurID:           r.CurID,
				Date:            dateParsed,
				CurAbbreviation: r.CurAbbreviation,
				CurScale:        r.CurScale,
				CurName:         r.CurName,
				CurOfficialRate: r.CurOfficialRate,
			})
		}

		err = repository.InsertCurrencyRates(db, toInsert)
		if err != nil {
			log.Println("Error inserting rates:", err)
			return
		}

		log.Println("Rates successfully inserted for date:", time.Now())
	})
	c.Start()
}
