package server

import (
	"groupie/server/model"
	"html/template"
	"net/http"
)

func GetRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errHandlers(w, r, http.StatusNotFound)
		return
	}
	if Tpl == nil {
		errHandlers(w, r, http.StatusInternalServerError)
		return
	}
	Tpl.Execute(w, artists)
}

func GetArtistById(w http.ResponseWriter, r *http.Request) {
	Tpl, err := template.ParseFiles("server/template/band.html")
	if err != nil {
		errHandlers(w, r, http.StatusInternalServerError)
		return

	}

	name := r.URL.Query().Get("name")
	var res model.Band
	if name == "" {
		errHandlers(w, r, http.StatusBadRequest)
		return
	} else {
		for i, v := range artists {
			if v.Name == name {
				res = artists[i]
				break
			}
		}
	}
	if res.Id == 0 {
		errHandlers(w, r, http.StatusNotFound)
		return
	}
	Tpl.Execute(w, res)
}
