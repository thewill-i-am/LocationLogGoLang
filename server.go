package main

import (
	routes "./controller"
	"./http"
	"fmt"
	"net/http"
)

var (
	httpRouter  = router.NewMuxRouter()
)

func main()  {
	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Running on port" + port)
	})
	httpRouter.GET("/posts", routes.GetPost)
	httpRouter.POST("/posts", routes.AddPost)
	httpRouter.SERVE(port)
}

