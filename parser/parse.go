package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// A Ar struct to map the Entire Response

// func main() {
// 	Locations("5", 5)
// }

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

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

var Artists []Artist

func Parser() []Artist {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &Artists) // unmarshalling our JSON into an array of Artist objects
	if err != nil {
		log.Fatal(err)
	}
	return Artists
}

// func Locations(id string, idNum int) {
// 	url := "https://groupietrackers.herokuapp.com/api/relation/"
// 	resp, err := http.Get(url + id)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var tempData Relations
// 	err = json.Unmarshal(body, &tempData) // unmarshalling our JSON into an array of Artist objects
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(tempData)
// }

func Locations(id string, idNum int) {
	url := "https://groupietrackers.herokuapp.com/api/relation/" + id
	tempReader, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempReader.Body.Close()
	body, err := ioutil.ReadAll(tempReader.Body)
	if err != nil {
		fmt.Println("err")
		return
	}
	var temp Relations
	if err = json.Unmarshal(body, &temp); err != nil {
		fmt.Println("err")
		return
	}

	Artists[idNum-1].DatesLocations = temp.DatesLocations
}
