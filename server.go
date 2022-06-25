package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"github.com/gorilla/mux"
)

func getDog(slc *[]string, wg *sync.WaitGroup) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal("Errors: ", err)
	}
	json, _ := ioutil.ReadAll(resp.Body)
	dog := string(json)
	*slc = append(*slc, dog)
	wg.Done()
}

func getDogNoWait(slc *[]string) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal("Errors: ", err)
	}
	json, _ := ioutil.ReadAll(resp.Body)
	dog := string(json)
	*slc = append(*slc, dog)
}

func getDogsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var wg sync.WaitGroup
	count, err := strconv.Atoi(vars["count"])
	if err != nil {
		log.Fatal("Errors: ", err)
	}

	wg.Add(count)
	var mySlice []string
	for i := 0; i < count; i++ {
		go getDog(&mySlice, &wg)
		// getDogNoWait(&mySlice)
	}
	wg.Wait()
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
