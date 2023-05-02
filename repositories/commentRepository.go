package repositories

import (
	"GoCommunityAPI/database"
	"GoCommunityAPI/database/entities"
	"GoCommunityAPI/models"
	"time"
)

var FetchCommentPageByArticleId func(articleId int, page int, pageSize int) (comments []models.CommentModel, err error)
var FindOneComment func(commentId int) (comment models.CommentModel, err error)
var CreateComment func(comment models.CommentModel) error
var DeleteComment func(commentId int) error

func init() {
	FetchCommentPageByArticleId = fetchCommentPageByArticleId
	FindOneComment = findOneComment
	CreateComment = createComment
	DeleteComment = deleteComment
}

func fetchCommentPageByArticleId(articleId int, page int, pageSize int) (comments []models.CommentModel, err error) {
	db := database.DB
	var entities []entities.CommentEntity
	var count int64
	transaction := db.Begin()
	db.Model(&entities).Count(&count)
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&entities)
	transaction.Model(&entities).
		Preload("User").
		Where("article_id = ?", articleId).
		Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&entities)
	err = transaction.Commit().Error
	comments = make([]models.CommentModel, len(entities))
	for index, entity := range entities {
		comments[index] = models.CommentModel{
			Id:        int(entity.ID),
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
	return comments, err
}

func findOneComment(commentId int) (comment models.CommentModel, err error) {
	db := database.DB
	var entity entities.CommentEntity
	query := db.Preload("User").Model(&entities.CommentEntity{})
	err = query.First(&entity, "id = ?", commentId).Error
	comment = models.CommentModel{
		Id:        int(entity.ID),
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
	return comment, err
}

func createComment(comment models.CommentModel) error {
	db := database.DB
	entity := entities.CommentEntity{
		Content:   comment.Content,
		UserId:    comment.UserId,
		ArticleId: comment.ArticleId,
	}
	err := db.Create(&entity).Error
	return err
}

func deleteComment(commentId int) error {
	db := database.DB
	query := db.Model(&entities.CommentEntity{}).Where("id = ?", commentId)
	err := query.Update("deleted_at", time.Time.Unix(time.Now())).Error
	return err
}
