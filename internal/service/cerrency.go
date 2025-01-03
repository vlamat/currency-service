package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NbrbRate struct {
	CurID           int     `json:"Cur_ID"`
	Date            string  `json:"Date"` // "2023-01-02T00:00:00"
	CurAbbreviation string  `json:"Cur_Abbreviation"`
	CurScale        int     `json:"Cur_Scale"`
	CurName         string  `json:"Cur_Name"`
	CurOfficialRate float64 `json:"Cur_OfficialRate"`
}

// FetchRatesFromNBRB получает курсы валют с сайта НБРБ
func FetchRatesFromNBRB() ([]NbrbRate, error) {
	url := "https://api.nbrb.by/exrates/rates?periodicity=0"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rates []NbrbRate
	if err := json.Unmarshal(body, &rates); err != nil {
		return nil, err
	}

	return rates, nil
}
