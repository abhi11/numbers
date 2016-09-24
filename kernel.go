package main

import (
	"log"
	"net/http"

	"github.com/abhi11/numbers/app/controllers"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("write something"))
}

func main() {
	http.HandleFunc("/", controllers.HelloWorld)
	http.HandleFunc("/numbers", controllers.SortNumbers)

	log.Println("Starting server on port 8888")
	http.ListenAndServe(":8888", nil)
}
