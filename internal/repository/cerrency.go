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

func InsertCurrencyRates(db *sql.DB, rates []CurrencyRate) error {
	// Чтобы сохранить все сразу
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	query := `
        INSERT INTO currency_rates (
            cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate
        ) VALUES (?, ?, ?, ?, ?, ?)
    `
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, rate := range rates {
		_, err := stmt.Exec(
			rate.CurID,
			rate.Date,
			rate.CurAbbreviation,
			rate.CurScale,
			rate.CurName,
			rate.CurOfficialRate,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// Получить все записи:
func GetAllRates(db *sql.DB) ([]CurrencyRate, error) {
	query := `
        SELECT id, cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate
        FROM currency_rates
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []CurrencyRate
	for rows.Next() {
		var r CurrencyRate
		err := rows.Scan(
			&r.ID,
			&r.CurID,
			&r.Date,
			&r.CurAbbreviation,
			&r.CurScale,
			&r.CurName,
			&r.CurOfficialRate,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}

// Получить записи за выбранный день:
func GetRatesByDate(db *sql.DB, date time.Time) ([]CurrencyRate, error) {
	query := `
        SELECT id, cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate
        FROM currency_rates
        WHERE date = ?
    `
	rows, err := db.Query(query, date.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []CurrencyRate
	for rows.Next() {
		var r CurrencyRate
		err := rows.Scan(
			&r.ID,
			&r.CurID,
			&r.Date,
			&r.CurAbbreviation,
			&r.CurScale,
			&r.CurName,
			&r.CurOfficialRate,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}
