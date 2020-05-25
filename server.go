package main

import (
	"./controller"
	"./http"
	"./repository"
	"./service"
	"fmt"
	"net/http"
)

var (
	postRepository = repository.NewFireStoreRepository()
	postService = service.NewPostService(postRepository)
	postController = controller.NewPostController(postService)
	httpRouterMux  = router.NewMuxRouter()
	httpRouterChi = router.NewChiRouter()
)

const port string = ":8080"

func main()  {
	httpRouterChi.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Running on port" + port)
	})
	httpRouterChi.GET("/posts", postController.GetPost)
	httpRouterChi.POST("/posts", postController.AddPost)
	httpRouterChi.SERVE(port)
}

