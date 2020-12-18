package main

import (
	"net/http"
	"github.com/IonutCraciun/filestorage/handlers"
	"log"
)


func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/view", handlers.ViewFile)
	http.HandleFunc("/update", handlers.UpdateFile)
	http.HandleFunc("/newfile", handlers.NewFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}