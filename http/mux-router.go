package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct {}

var(
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request))  {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request))  {
	 muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) SERVE(port string){
	fmt.Printf("Listening on port" + port)
	http.ListenAndServe(port, muxDispatcher)
}