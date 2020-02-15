package main

import (
	"comical/feed/handlers"
	"comical/model/superfeedr"
	"comical/rest"
	"comical/storage"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)
//import "google.golang.org/appengine/log"

func main() {
	http.Handle("/", rest.AdaptFunc(indexHandler, rest.Log()))
	http.Handle("/feed/test/update", rest.AdaptFunc(testUpdateHandler, rest.Log(), storage.Firestore()))
	http.Handle("/feed/update", rest.AdaptFunc(updateHandler, rest.Log(), storage.Firestore()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Println(os.Getenv("FIREBASE_CRED"))

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
	if r.URL.Path != "/feed/update" {
		http.NotFound(w, r)
		return
	}

	//client := storage.GetClient()

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

	for _, item := range feedItem.Items {
		entry, err := handlers.ExtractEntry(&item)

		if err != nil {
			log.Printf("unable to extract entry: %v", err)
		} else {
			log.Printf("entry: %v", entry)

			err = storage.NewComicsEntry(r.Context(), entry)
			if err != nil {
				log.Printf("unable to write to storage: %v", err)
			}
		}
	}

	log.Printf("Update: %v", feedItem)

	_, err := fmt.Fprint(w, "Done updating")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func testUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/feed/test/update" {
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