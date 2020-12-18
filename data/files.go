package data

import (
	"io/ioutil"
	"log"
)

// File struct
type File struct {
	Title string
	Body  []byte
	Cookie string
}

// GetAllFileNames function public
func GetAllFileNames() (titles []string) {
	files, err := ioutil.ReadDir("files")
	if err != nil {
    	log.Fatal(err)
	}
	for _ , val := range files {
		titles = append(titles, val.Name())
	}
	return titles
}