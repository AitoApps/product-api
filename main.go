package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	port := ":3000"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Data %s", d)
	})

	http.ListenAndServe(port, nil)
}