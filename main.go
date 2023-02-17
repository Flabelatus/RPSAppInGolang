package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"myApp/rps"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	// struct of rps.Round
	result := rps.PlayRound(playerChoice)

	// convert to JSON
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func main() {

	// run the homepage function on the route "/"
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)

	fmt.Println("Starting the webserver on port 8080")
	// running the application on the http://localhost:8080
	http.ListenAndServe(":8080", nil)

}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
