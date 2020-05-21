package routes

import (
	"../entity"
	"../repository"
	"encoding/json"
	"math/rand"
	"net/http"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func GetPost(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error":"Error getting data"}`))
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}
func AddPost(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewEncoder(res).Encode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error":"Error parsing data"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)
}

