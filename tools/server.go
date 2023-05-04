package tools

import (
	"log"
	"net/http"
)

func Run() {
	file := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", file))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
