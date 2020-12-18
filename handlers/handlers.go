package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/IonutCraciun/filestorage/data"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	log.Printf("handlers: %s %s %s ", r.Method, r.URL.Path, r.Proto)

	filesName := data.GetAllFileNames()
	log.Printf("%+v", filesName)
	err := tpl.ExecuteTemplate(w, "index.html", filesName)
	if err != nil {
		log.Fatalln(err)
	}

}

// ViewFile display file
func ViewFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("handlers: %s %s %s ", r.Method, r.URL.Path, r.Proto)
	fileName := r.FormValue("filename")
	if len(fileName) == 0 {
		http.Error(w, "File name parameter not provided", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadFile(fmt.Sprintf("files/%v", fileName))
	if err != nil {
		http.Error(w, fmt.Sprintf("File %v don't exists on the server", fileName), http.StatusBadRequest)
		return
	}
	f := data.File{Title: fileName, Body: body, Cookie: ""}
	tpl.ExecuteTemplate(w, "viewFile.html", f)
}

// UpdateFile save
func UpdateFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST methods accepted", http.StatusBadRequest)
	}
	log.Printf("handlers: %s %s %s ", r.Method, r.URL.Path, r.Proto)
	fileName := r.FormValue("filename")
	if len(fileName) == 0 {
		http.Error(w, "File name parameter not provided", http.StatusBadRequest)
		return
	}

	body := r.FormValue("body")
	err := ioutil.WriteFile((fmt.Sprintf("files/%v", fileName)), []byte(body), 0644)
	if err != nil {
		http.Error(w, fmt.Sprintf("Server error updating file %v", fileName), http.StatusInternalServerError)
		return
	}
	f := data.File{Title: fileName, Body: []byte(body), Cookie: ""}
	tpl.ExecuteTemplate(w, "fileUpdated.html", f)

}

// NewFile create a new file
func NewFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("handlers: %s %s %s ", r.Method, r.URL.Path, r.Proto)
	switch r.Method {
	case "GET":
		{
			tpl.ExecuteTemplate(w, "newFile.html", nil)
		}
	case "POST":
		{
			fileName := r.FormValue("filename")
			if _, err := os.Stat(fmt.Sprintf("files/%v", fileName)); err == nil {
				http.Error(w, fmt.Sprintf("File  '%v' already exists on the server", fileName), http.StatusBadRequest)
				return
			}
			body := r.FormValue("body")
			err := ioutil.WriteFile((fmt.Sprintf("files/%v", fileName)), []byte(body), 0644)
			if err != nil {
				http.Error(w, fmt.Sprintf("Server error creating file %v", fileName), http.StatusInternalServerError)
				return
			}
			tpl.ExecuteTemplate(w, "fileCreated.html", nil)
		}
	default:
		http.Error(w, "Server error. Method unknown", http.StatusInternalServerError)
		return
	}
}

// DeleteFile delete a file
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	// e.g: curl -X DELETE localhost:8080/delete?filename=test -v or curl --location --request DELETE 'localhost:8080/delete?filename=123456'
	log.Printf("handlers: %s %s %s ", r.Method, r.URL.Path, r.Proto)
	if r.Method != "DELETE" {
		http.Error(w, "Only DELETE method accepted", http.StatusBadRequest)
	}
	filename := r.FormValue("filename")
	if _, err := os.Stat(fmt.Sprintf("files/%v", filename)); err != nil {
		http.Error(w, fmt.Sprintf("File '%v' doesn't exist on the server", filename), http.StatusBadRequest)
		return
	}
	err := os.Remove(fmt.Sprintf("files/%v", filename))
	if err != nil {
		http.Error(w, "Server Error. Couldn't delete the file", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("\n File %v deleted succesfully \n", filename))
	log.Printf("handlers: File %v deleted succesfully", filename)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
