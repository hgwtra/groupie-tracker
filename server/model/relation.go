package model

//type Map map[string]string
type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Relations []Relation `json:"index"`
}

type NewRelation struct {
	Location string
	Dates    []string
}
