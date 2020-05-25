package repository

import (
	entity "../entity"
	enviroment "../enviroment"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
)

type repo struct {}

func NewFireStoreRepository() PostRepository{
	return &repo{}
}

const (
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile(enviroment.GoDotEnvVariable("credentials"))
	app, err := firebase.NewApp(ctx, nil, sa)
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Fail to create a FireStone Client: %v", err)
		return nil, err
	}
	defer client.Close()
	_,_,err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil{
		log.Fatalf("Fail Adding a New Post: %v", err)
		return nil, err
	}
	return post,nil
}

func (*repo) FindAll() ([]entity.Post, error){
	ctx := context.Background()
	sa := option.WithCredentialsFile(enviroment.GoDotEnvVariable("credentials"))
	app, err := firebase.NewApp(ctx, nil, sa)
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Fail to create a FireStone Client: %v", err)
		return nil, err
	}
	if err != nil {
		log.Fatalf("Fail to create a FireStone Client: %v", err)
		return nil, err
	}
	defer client.Close()
	iter := client.Collection(collectionName).Documents(ctx)
	var posts []entity.Post
	for  {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err !=  nil{
			log.Fatalf("Fail to iterate list of post: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID: doc.Data()["ID"].(int64),
			Text: doc.Data()["Text"].(string),
			Title: doc.Data()["Title"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}