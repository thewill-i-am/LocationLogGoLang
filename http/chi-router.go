package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type chaiRouter struct {}

var(
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router  {
	return &chaiRouter{}
}

func(*chaiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)){
	chiDispatcher.Get(uri, f)
}

func(*chaiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)){
	chiDispatcher.Post(uri, f)
}

func(*chaiRouter) SERVE(port string){
	fmt.Printf("Listening on port " + port)
	http.ListenAndServe(port, chiDispatcher)
}


