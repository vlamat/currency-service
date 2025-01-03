package main

import (
	"log"

	"currency/internal/db"
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
}
