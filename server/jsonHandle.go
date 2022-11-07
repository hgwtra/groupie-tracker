package server

import (
	"encoding/json"
	"groupie/server/model"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

func getAPI(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}

func handleJSON() []model.Band {
	var res []model.Band
	//get artists
	artistsData := getAPI("https://groupietrackers.herokuapp.com/api/artists")
	json.Unmarshal([]byte(artistsData), &res)
	//get locations
	var locations model.Locations
	locationsData := getAPI("https://groupietrackers.herokuapp.com/api/locations")
	json.Unmarshal([]byte(locationsData), &locations)
	for i, v := range locations.Locations {
		res[i].Locations = v.Location
	}
	//get dates
	var dates model.Dates
	datesData := getAPI("https://groupietrackers.herokuapp.com/api/dates")
	json.Unmarshal([]byte(datesData), &dates)
	for i, v := range dates.Dates {
		res[i].ConcertDates = v.Date
	}
	//get relations
	var relations model.Relations
	relationsData := getAPI("https://groupietrackers.herokuapp.com/api/relation")
	json.Unmarshal([]byte(relationsData), &relations)
	for i, v := range relations.Relations {
		for x, y := range v.DatesLocations {
			var newRelation model.NewRelation
			newRelation.Location = x
			newRelation.Dates = y

			res[i].Relations = append(res[i].Relations, newRelation)
		}
	}
	// replace/change data
	for i, v := range res {
		//sort members
		sort.Strings(v.Members)
		//change first album date(-) to date(.)
		res[i].FirstAlbum = strings.ReplaceAll(v.FirstAlbum, "-", ".")
		//Methods for imported type​​ Methods can be defined only inside the package where type is created.
		//Locations: sort, change (-) to ( ), (_)  to (, ), uppercase
		sort.Strings(v.Locations)
		for i := range v.Locations {
			v.Locations[i] = strings.ReplaceAll(v.Locations[i], "_", " ")
			v.Locations[i] = strings.ReplaceAll(v.Locations[i], "-", ", ")
			v.Locations[i] = strings.Title(v.Locations[i])
		}
		//concertDates: remove *, change (-) to (.)
		for i := range v.ConcertDates {
			v.ConcertDates[i] = strings.ReplaceAll(v.ConcertDates[i], "*", "")
			v.ConcertDates[i] = strings.ReplaceAll(v.ConcertDates[i], "-", ".")
		}
		//relations- changeg location to (-) (""), date(-) to (.) format,
		var temp []string
		for i := range v.Relations {
			v.Relations[i].Location = strings.ReplaceAll(v.Relations[i].Location, "_", " ")
			v.Relations[i].Location = strings.ReplaceAll(v.Relations[i].Location, "-", ", ")
			v.Relations[i].Location = strings.Title(v.Relations[i].Location)
			for j := range v.Relations[i].Dates {
				v.Relations[i].Dates[j] = strings.ReplaceAll(v.Relations[i].Dates[j], "-", ".")
				if v.Relations[i].Dates[j] == v.ConcertDates[0] {
					temp = append(temp, v.Relations[i].Location)
				}
			}

		}

		res[i].RecentConcerts = temp

	}
	return res
}
