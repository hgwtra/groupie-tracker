package model

type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

type Locations struct {
	Locations []Location `json:"index"`
}
