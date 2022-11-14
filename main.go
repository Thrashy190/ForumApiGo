package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("connecting to 3000")

	r := mux.NewRouter();

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Connection successful on port 3000"))
	});

	http.ListenAndServe(":3000",r);
}