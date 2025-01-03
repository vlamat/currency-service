package repository

import (
	"database/sql"
	"time"
)

type CurrencyRate struct {
	ID              int
	CurID           int
	Date            time.Time
	CurAbbreviation string
	CurScale        int
	CurName         string
	CurOfficialRate float64
}

func InsertCurrencyRate(db *sql.DB, rate CurrencyRate) error {
	query := `
        INSERT INTO currency_rates (
            cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate
        ) VALUES (?, ?, ?, ?, ?, ?)
    `
	_, err := db.Exec(query,
		rate.CurID,
		rate.Date,
		rate.CurAbbreviation,
		rate.CurScale,
		rate.CurName,
		rate.CurOfficialRate,
	)
	return err
}
