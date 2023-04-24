package services

import (
	"GoCommunityAPI/models"
	"GoCommunityAPI/repositories"
)

func GetArticleList(page int, pageSize int) ([]models.ArticleModel, error) {
	articles, err := repositories.FetchArticlesPage(page, pageSize)
	return articles, err
}

func GetArticleDetail(articleId int) (models.ArticleModel, error) {
	article, err := repositories.FindOneArticle(articleId)
	return article, err
}

func UploadArticle(article models.ArticleModel) error {
	err := repositories.CreateArticle(article)
	return err
}

func UpdateArticle(article models.ArticleModel) error {
	err := repositories.UpdateArticle(article)
	return err
}

func DeleteArticle(articleId int) error {
	err := repositories.DeleteArticle(articleId)
	return err
}
