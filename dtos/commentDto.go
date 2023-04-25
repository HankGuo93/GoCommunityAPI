package dtos

import "GoCommunityAPI/models"

type CommentDto struct {
	Id        int
	Content   string
	UserId    int
	ArticleId int
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64

	User UserDto
}

func CreateCommentPageResponse(comments []models.CommentModel) []map[string]interface{} {
	var dtos = make([]map[string]interface{}, len(comments))
	for index, model := range comments {
		dtos[index] = map[string]interface{}{
			"Id":        model.Id,
			"ArticleId": model.ArticleId,
			"Content":   model.Content,
			"CreatedAt": model.CreatedAt,
			"UpdatedAt": model.UpdatedAt,
			"User": map[string]interface{}{
				"Id":    model.User.Id,
				"Name":  model.User.Name,
				"Email": model.User.Email,
			},
		}
	}
	return dtos
}
