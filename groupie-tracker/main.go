package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Group struct {
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
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fs = http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	fs = http.FileServer(http.Dir("scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fs))
	http.HandleFunc("/Groups.html", func(w http.ResponseWriter, r *http.Request) {
		response, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		data, _ := ioutil.ReadAll(response.Body)
		bytes := []byte(data)
		var list []Group
		json.Unmarshal(bytes, &list)
		tmpl, _ := template.ParseFiles("./html/Groups.html")
		tmpl.Execute(w, list)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path != "/relation" {
			tmpl, _ := template.ParseFiles("./html/ERROR.html")
			tmpl.Execute(w, nil)
			return
		}
		response, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		data, _ := ioutil.ReadAll(response.Body)
		bytes := []byte(data)
		var list []Group
		json.Unmarshal(bytes, &list)
		tmpl, _ := template.ParseFiles("./html/index.html")
		tmpl.Execute(w, list)
	})
	fmt.Println("Server is listening ... http://localhost:8181/")
	http.ListenAndServe(":8181", nil)

}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	// A slice of sample websites
// 	urls := []string{
// 		"https://www.easyjet.com/",
// 		"https://www.skyscanner.de/",
// 		"https://www.ryanair.com",
// 		"https://wizzair.com/",
// 		"https://www.swiss.com/",
// 	}

// 	c := make(chan urlStatus)
// 	for _, url := range urls {
// 		go checkUrl(url, c)

// 	}
// 	result := make([]urlStatus, len(urls))
// 	for i, _ := range result {
// 		result[i] = <-c
// 		if result[i].status {
// 			fmt.Println(result[i].url, "is up.")
// 		} else {
// 			fmt.Println(result[i].url, "is down !!")
// 		}
// 	}

// }

// //checks and prints a message if a website is up or down
// func checkUrl(url string, c chan urlStatus) {
// 	_, err := http.Get(url)
// 	if err != nil {
// 		// The website is down
// 		c <- urlStatus{url, false}
// 	} else {
// 		// The website is up
// 		c <- urlStatus{url, true}
// 	}
// }

// type urlStatus struct {
// 	url    string
// 	status bool
// }
