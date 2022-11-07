package server

import (
	"html/template"
	"log"
	"net/http"
)

type errorData struct {
	Num  int
	Text string
}

func errHandlers(w http.ResponseWriter, r *http.Request, err int) {
	temp, er := template.ParseFiles("server/template/error.html")
	if er != nil {
		log.Fatal(er)
		return
	}
	w.WriteHeader(err)
	errData := errorData{Num: err}
	if err == 404 {
		errData.Text = "Page Not Found"
	} else if err == 400 {
		errData.Text = "Bad Request"
	} else if err == 500 {
		errData.Text = "Internal Server Error"
	}
	temp.Execute(w, errData)
}
