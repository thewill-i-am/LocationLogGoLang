package main

import (
	routes "./routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		fmt.Fprintln(res, "Running")
	})
	router.HandleFunc("/posts", routes.GetPost).Methods("GET")
	router.HandleFunc("/posts", routes.AddPost).Methods("POST")
	log.Println("Server on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

