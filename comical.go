package main

import (
	"comical/storage"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"comical/model/superfeedr"
)
//import "google.golang.org/appengine/log"

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/feed/update", updateHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	client, storageErr := storage.InitStorage()

	if storageErr != nil {
		log.Printf("Unable to create firebase cloud store: %v", storageErr)
	} else {
		log.Printf("Client created: %v", client)
	}

	if r.URL.Path != "/feed/update" {
		http.NotFound(w, r)
		return
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	feedItem := superfeedr.FeedUpdate{}
	log.Print(string(bodyBytes))

	parseError := json.Unmarshal(bodyBytes, &feedItem)
	if parseError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "Error: %v", parseError)
		log.Printf("Error: %v", parseError)
		return
	}
	log.Printf("Update: %v", feedItem)

	_, err := fmt.Fprint(w, "Done updating")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}