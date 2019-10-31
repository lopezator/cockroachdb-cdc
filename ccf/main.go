package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func cdcHandler(w http.ResponseWriter, r *http.Request) {
	// get request body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// print body request to stdout
	fmt.Printf("%s\n", b)
}

func main() {
	http.HandleFunc("/cdc/", cdcHandler)
	fmt.Println("started file server on localhost:8080/cdc")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}