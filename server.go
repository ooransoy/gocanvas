package main

import (
	"html/template"
	"image/png"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func serve() {
	resetCanvas()
	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/img", imgHandler).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("static")),
	))

	log.Fatal(http.ListenAndServe(":3000", router))
}

func handler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	tpl, err := template.ParseFiles("index.tpl")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(writer, "Walker")
}

func imgHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "image/png")
	err := png.Encode(writer, canvas)
	if err != nil {
		if err.Error()[:9] == "write tcp" {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}
}
