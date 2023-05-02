package repositories

import (
	"GoCommunityAPI/database"
	"GoCommunityAPI/database/entities"
	"GoCommunityAPI/models"
	"strings"
	"time"
)

var FetchArticlesPage func(page int, pageSize int) (article []models.ArticleModel, err error)
var FindOneArticle func(articleId int) (article models.ArticleModel, err error)
var CreateArticle func(article models.ArticleModel) error
var UpdateArticle func(article models.ArticleModel) error
var DeleteArticle func(articleId int) error

func init() {
	FetchArticlesPage = fetchArticlesPage
	FindOneArticle = findOneArticle
	CreateArticle = createArticle
	UpdateArticle = updateArticle
	DeleteArticle = deleteArticle
}

func fetchArticlesPage(page int, pageSize int) (article []models.ArticleModel, err error) {
	db := database.DB
	var entities []entities.ArticleEntity
	var count int64
	transaction := db.Begin()
	db.Model(&entities).Count(&count)
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&entities)
	transaction.Model(&entities).
		Preload("User").
		Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&entities)
	err = transaction.Commit().Error
	article = make([]models.ArticleModel, len(entities))
	for index, entity := range entities {
		article[index] = models.ArticleModel{
			Id:        int(entity.ID),
			Title:     entity.Title,
			Content:   entity.Content,
			UserId:    entity.UserId,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,

			User: models.UserModel{
				Id:        int(entity.User.ID),
				Name:      entity.User.Name,
				Email:     entity.User.Email,
				Password:  entity.User.Password,
				CreatedAt: entity.User.CreatedAt,
				UpdatedAt: entity.User.UpdatedAt,
			},
		}
	}
	return article, err
}

func findOneArticle(articleId int) (article models.ArticleModel, err error) {
	db := database.DB
	var entity entities.ArticleEntity
	query := db.Preload("User").Model(&entities.ArticleEntity{})
	err = query.First(&entity, "id = ?", articleId).Error
	article = models.ArticleModel{
		Id:        int(entity.ID),
		Title:     entity.Title,
		Content:   entity.Content,
		UserId:    entity.UserId,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,

		User: models.UserModel{
			Id:        int(entity.User.ID),
			Name:      entity.User.Name,
			Email:     entity.User.Email,
			Password:  entity.User.Password,
			CreatedAt: entity.User.CreatedAt,
			UpdatedAt: entity.User.UpdatedAt,
		},
	}
	return article, err
}

func createArticle(article models.ArticleModel) error {
	db := database.DB
	entity := entities.ArticleEntity{
		Title:   article.Title,
		Content: article.Content,
		UserId:  article.UserId,
	}
	err := db.Create(&entity).Error
	return err
}

func updateArticle(article models.ArticleModel) error {
	db := database.DB
	updateMapping := make(map[string]interface{})
	entity := entities.ArticleEntity{
		Title:   article.Title,
		Content: article.Content,
	}
	entity.ID = uint(article.Id)
	if len(strings.TrimSpace(entity.Title)) != 0 {
		updateMapping["Title"] = entity.Title
	}
	if len(strings.TrimSpace(entity.Content)) != 0 {
		updateMapping["Content"] = entity.Content
	}
	err := db.Model(entity).Updates(updateMapping).Error
	return err
}

func deleteArticle(articleId int) error {
	db := database.DB
	query := db.Model(&entities.ArticleEntity{}).Where("id = ?", articleId)
	err := query.Update("deleted_at", time.Time.Unix(time.Now())).Error
	return err
}
