package services

import (
	"GoCommunityAPI/models"
	"GoCommunityAPI/repositories"
	"GoCommunityAPI/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadComment_Success(t *testing.T) {
	//arrange
	comment := models.CommentModel{
		UserId:    1,
		ArticleId: 1,
		Content:   "Test Comment",
	}
	repositories.FindOneUserById = func(userId int) (models.UserModel, error) {
		return models.UserModel{}, nil
	}
	repositories.FindOneArticle = func(articleId int) (models.ArticleModel, error) {
		return models.ArticleModel{}, nil
	}
	repositories.CreateComment = func(comment models.CommentModel) error {
		return nil
	}

	//act
	err := services.UploadComment(comment)

	//assert
	assert.Nil(t, err)
}

func TestUploadComment_InvalidUserId(t *testing.T) {
	//arrange
	comment := models.CommentModel{
		UserId:    1,
		ArticleId: 1,
		Content:   "Test Comment",
	}
	repositories.FindOneUserById = func(userId int) (models.UserModel, error) {
		return models.UserModel{}, errors.New("User not found")
	}

	//act
	err := services.UploadComment(comment)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, "You must provide a valid user id", err.Error())
}

func TestUploadComment_InvalidArticleId(t *testing.T) {
	//arrange
	comment := models.CommentModel{
		UserId:    1,
		ArticleId: 1,
		Content:   "Test Comment",
	}
	repositories.FindOneUserById = func(userId int) (models.UserModel, error) {
		return models.UserModel{}, nil
	}
	repositories.FindOneArticle = func(articleId int) (models.ArticleModel, error) {
		return models.ArticleModel{}, errors.New("Article not found")
	}

	//act
	err := services.UploadComment(comment)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, "You must provide a valid article id", err.Error())
}

func TestDeleteComment_Success(t *testing.T) {
	//arrange
	commentId := 1
	repositories.FindOneComment = func(commentId int) (models.CommentModel, error) {
		return models.CommentModel{}, nil
	}
	repositories.DeleteComment = func(commentId int) error {
		return nil
	}

	//act
	err := services.DeleteComment(commentId)

	//assert
	assert.Nil(t, err)
}

func TestDeleteComment_InvalidCommentId(t *testing.T) {
	//arrange
	commentId := 1
	repositories.FindOneComment = func(commentId int) (models.CommentModel, error) {
		return models.CommentModel{}, errors.New("Comment not found")
	}

	//act
	err := services.DeleteComment(commentId)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, "You must provide a valid comment id", err.Error())
}
