package service

import (
	entity "../entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateEmpty(t *testing.T)  {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post is empty")
}
func TestValidateEmptyPostTitle(t *testing.T)  {
	post := entity.Post{ID: 1, Title: "", Text: "B"}
	testService := NewPostService(nil)
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Post title is empty")
}