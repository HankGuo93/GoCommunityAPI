package repositories

import (
	"GoCommunityAPI/models"
	"errors"
)

func FetchArticlesPage(page int, pageSize int) ([]models.ArticleModel, error) {
	return []models.ArticleModel{}, errors.New("Not Implemented")
}

func FindOneArticle(articleId int) (models.ArticleModel, error) {
	return models.ArticleModel{}, errors.New("Not Implemented")
}

func CreateArticle(article models.ArticleModel) error {
	return errors.New("Not Implemented")
}

func UpdateArticle(article models.ArticleModel) error {
	return errors.New("Not Implemented")
}

func DeleteArticle(articleId int) error {
	return errors.New("Not Implemented")
}
