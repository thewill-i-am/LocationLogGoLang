package repository
import (
	entity "../entity"
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)

}

type repo struct {}

func NewPostRepository() PostRepository{
	return &repo{}
}

const (
	projectId string = "goprueba-2c638"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
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
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Fail to create a FireStone Client: %v", err)
		return nil, err
	}
	defer client.Close()
	iterator := client.Collection(collectionName).Documents(ctx)
	var posts []entity.Post
	for  {
		doc, err := iterator.Next()
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