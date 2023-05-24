package handlers

import (
	"fmt"
	parser "groupie-tracker/parser"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

type Artist struct {
	Id             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstalbum"`
	Relations      string   `json:"relations"`
	DatesLocations map[string][]string
}

func ErrExec(r http.ResponseWriter, header int) {
	r.WriteHeader(header)
	tempErr, err := template.ParseFiles("./template/error.html")
	if err != nil {
		ErrExec(r, http.StatusNotFound)
		log.Print(err)
		return
	}
	tempErr.ExecuteTemplate(r, "error", header)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrExec(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrExec(w, http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("./template/index.html")
	if err != nil {
		ErrExec(w, http.StatusNotFound)
		log.Print(err)
		return
	}

	Info := parser.Parser()
	err = temp.ExecuteTemplate(w, "index", Info)
	if err != nil {
		ErrExec(w, http.StatusNotFound)
		log.Print(err)
		return
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/artist/" {
	// ErrExec(w, http.StatusNotFound)
	// return
	// }
	if r.Method != http.MethodGet {
		ErrExec(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi((path.Base(r.URL.Path)))
	if err != nil {
		fmt.Println(err)
		ErrExec(w, http.StatusNotFound)
		return
	}
	if id < 1 || id > 52 {
		ErrExec(w, http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("./template/artist.html")
	if err != nil {
		ErrExec(w, http.StatusNotFound)
		log.Print(err)
		return
	}
	Info := parser.Parser()
	parser.Locations(strconv.Itoa(id), id)
	err = temp.ExecuteTemplate(w, "artist", Info[id-1])
	if err != nil {
		log.Print(err)
		return
	}
}
