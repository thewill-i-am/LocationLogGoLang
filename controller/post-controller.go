package controller

import (
	entity "../entity"
	"../errors"
	"../service"
	"encoding/json"
	"net/http"
)

type controller struct {}

var (
	postService service.PostService
)

type PostController interface {
	GetPost(res http.ResponseWriter, req *http.Request)
	AddPost(res http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPost(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting post"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func (*controller) AddPost(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
	var post entity.Post
	err := decoder.Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error parsing post"})
		return
	}
	errValidation := postService.Validate(&post)
	if errValidation != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: errValidation.Error()})
		return
	}
	postService.Create(&post)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)
}

