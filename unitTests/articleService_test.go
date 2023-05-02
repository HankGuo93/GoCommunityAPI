package services_test

import (
	"GoCommunityAPI/models"
	"GoCommunityAPI/repositories"
	"GoCommunityAPI/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateArticle_Success(t *testing.T) {
	//arrange
	mockArticle := models.ArticleModel{Id: 1, Title: "Mock Article", Content: "This is a mock article content."}
	repositories.FindOneArticle = func(articleId int) (models.ArticleModel, error) {
		return mockArticle, nil
	}
	repositories.UpdateArticle = func(article models.ArticleModel) error {
		return nil
	}

	//action
	err := services.UpdateArticle(mockArticle)

	//assert
	if err != nil {
		t.Errorf("UpdateArticle() returned an error: %v", err)
	}
}

func TestUpdateArticle_Fail(t *testing.T) {
	//arrange
	mockArticle := models.ArticleModel{Id: 1, Title: "Mock Article", Content: "This is a mock article content."}
	repositories.FindOneArticle = func(articleId int) (models.ArticleModel, error) {
		return mockArticle, errors.New("not found")
	}
	repositories.UpdateArticle = func(article models.ArticleModel) error {
		return nil
	}

	//action
	err := services.UpdateArticle(mockArticle)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, "not found", err.Error())
}

func TestDeleteArticle_Success(t *testing.T) {
	//arrange
	mockArticleId := 1
	repositories.FindOneArticle = func(articleId int) (models.ArticleModel, error) {
		return models.ArticleModel{}, nil
	}
	repositories.DeleteArticle = func(articleId int) error {
		return nil
	}

	//action
	err := services.DeleteArticle(mockArticleId)

	//assert
	if err != nil {
		t.Errorf("DeleteArticle() returned an error: %v", err)
	}
}

func TestDeleteArticle_Fail(t *testing.T) {
	//arrange
	mockArticleId := 1
	repositories.FindOneArticle = func(articleId int) (models.ArticleModel, error) {
		return models.ArticleModel{}, errors.New("not found")
	}
	repositories.DeleteArticle = func(articleId int) error {
		return nil
	}

	//action
	err := services.DeleteArticle(mockArticleId)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, "not found", err.Error())
}
