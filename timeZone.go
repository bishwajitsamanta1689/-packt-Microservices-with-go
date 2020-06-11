package main

import (
	"fmt"
	"net/http"
)

func timezone(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	zone := r.URL.Query().Get("tz") // Query should be like := http://localhost:8000/?tz=IST
	_, err := fmt.Fprint(w, "Hello Current Time Zone::"+zone)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", timezone)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
