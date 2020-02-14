package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type ArtistData struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func main() {
	fmt.Println("Starting the application...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {

			response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				bytes := []byte(data)
				var Art []ArtistData
				json.Unmarshal(bytes, &Art)

				tmpl, _ := template.ParseFiles("index.html")
				if r.Method == "GET" {
					tmpl.Execute(w, Art)

				}
			}
		} else if r.URL.Path == "/relation" {

			tmpl, _ := template.ParseFiles("relation.html")
			if r.Method == "GET" {
				tmpl.Execute(w, nil)
			}
		}

	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}
