package model

type Date struct {
	Id   int      `json:"id"`
	Date []string `json:"dates"`
}

type Dates struct {
	Dates []Date `json:"index"`
}
