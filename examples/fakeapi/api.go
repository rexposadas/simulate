package fakeapi

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"strconv"
)

// StartAPI starts a fake api
func StartAPI(port int) {
	r := mux.NewRouter()
	r.HandleFunc("/jobs", A)
	http.Handle("/", r)

	fmt.Printf("Started fake API on PORT %d\n", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}
}

// A is a simple handler which prints request method
func A(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "%s", "received GET request")
	case "POST":
		fmt.Fprintf(w, "%s", "received POST request")
	}
}
