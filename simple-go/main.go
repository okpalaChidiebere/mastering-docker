package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func sleep(ms int){
	time.Sleep(time.Duration(ms)*time.Millisecond)
}
	
func main() {

	r := mux.NewRouter()

	//background thread task
	go func(){
		for {
			log.Printf("containers rule!")
			sleep(5000)
		}
	}()

    r.HandleFunc("/health", func (w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello!")
	}).Methods("GET")

    http.ListenAndServe(":8080", r)
}