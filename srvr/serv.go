package srvr

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Привет! Ты запросил: %d из папки: <span style='color:red'>%s</span></h1>", Varr(), r.URL.Path)
}

func Serv() {
	http.HandleFunc("/", handler)

	fmt.Println("Сервер запущен на http://localhost:8080") // message to console
	http.ListenAndServe(":8080", nil)
}
