package server

import (
	"groupie/server/model"
	"html/template"
	"log"
	"net/http"
)

//create a variable from server package
var Tpl *template.Template
var artists []model.Band

func Start() error {
	artists = handleJSON()
	tmpl, err := template.ParseFiles("server/template/index.html")
	if err != nil {
		//log.Fatal(err)
		tmpl = nil
	}
	Tpl = tmpl
	//handle css from static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./server/static/"))))
	//get request
	http.HandleFunc("/", GetRequest)
	http.HandleFunc("/artists", GetArtistById)
	//post request

	//open port- listen
	log.Println("Staring server on port...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
