package main

import (
	"log"
	"net/http"

	"github.com/IonutCraciun/filestorage/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/view", handlers.ViewFile)
	http.HandleFunc("/update", handlers.UpdateFile)
	http.HandleFunc("/newfile", handlers.NewFile)
	http.HandleFunc("/delete", handlers.DeleteFile) // accesible only by http request with curl, etc.
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
