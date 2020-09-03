package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Artist struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Members   string `json:"members"`
	Locations string `json:"locations"`
	Relations string `json:"relations"`
	Date 	  int	 `json:"creationDate"`	
}

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {

			templates := template.Must(template.ParseGlob("view/*.html"))

			artists := DecodeArrayMap("https://groupietrackers.herokuapp.com/api/artists")

			params := map[string]interface{}{"Artists": artists}

			templates.ExecuteTemplate(w, "templates.html", params)

			fmt.Println(params)

		}
	}
}
func DecodeArrayMap(url string) []Artist {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artists []Artist

	json.Unmarshal(data, &artists)
	return artists
}
