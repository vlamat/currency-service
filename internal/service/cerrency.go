package service

type NbrbRate struct {
	CurID           int     `json:"Cur_ID"`
	Date            string  `json:"Date"` // "2023-01-02T00:00:00"
	CurAbbreviation string  `json:"Cur_Abbreviation"`
	CurScale        int     `json:"Cur_Scale"`
	CurName         string  `json:"Cur_Name"`
	CurOfficialRate float64 `json:"Cur_OfficialRate"`
}
