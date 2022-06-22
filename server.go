package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

func getDog(slc *[]string, channel chan string) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal("Errors: ", err)
	}
	json, _ := ioutil.ReadAll(resp.Body)
	dog := string(json)
	channel <- dog
	*slc = append(*slc, <- channel)
}

func getDogsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	count, err := strconv.Atoi(vars["count"])
	if err != nil {
		log.Fatal("Errors: ", err)
	}

	var mySlice []string
	c := make(chan string, count)
	for i := 0; i < count; i++ {
		go getDog(&mySlice, c)
	}

	// Switch to sync.WaitGroup to increase stability
	time.Sleep(time.Second * 2)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mySlice)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/dogs/{count}", getDogsHandler).Methods("Get")
	staticDir := "/static/"
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Errors: ", err)
	}
}
