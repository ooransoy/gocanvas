package gocanvas

import (
	"encoding/base64"
	"bytes"
	"html/template"
	"image/png"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func serve() {
	ResetCanvas()
	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/img", imgHandler).Methods("GET")
	router.HandleFunc("/ws", wsHandler).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("static")),
	))

	log.Fatal(http.ListenAndServe(":3000", router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl, err := template.ParseFiles("index.tpl")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, map[string]string{
		"Title": "GoCanvas",
		"Ws": "ws://"+r.Host+"/ws",
	})
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	ticker := time.NewTicker(time.Second / 2)
	defer ticker.Stop()

	var buf bytes.Buffer
	for {
		<-ticker.C
		png.Encode(&buf, Canvas)
		b64str := base64.StdEncoding.EncodeToString(buf.Bytes())
		ws.WriteMessage(websocket.TextMessage, []byte(b64str))
	}
}

func imgHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "image/png")
	err := png.Encode(writer, Canvas)
	if err != nil {
		if err.Error()[:9] == "write tcp" {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}
}
