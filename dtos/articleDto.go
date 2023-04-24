package dtos

import "GoCommunityAPI/models"

type ArticleDto struct {
	Id        int
	Title     string
	Content   string
	UserId    int
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64

	User UserDto
}

func CreateArticlePageResponse(articles []models.ArticleModel) []map[string]interface{} {
	var dtos = make([]map[string]interface{}, len(articles))
	for index, model := range articles {
		dtos[index] = map[string]interface{}{
			"Id":        model.Id,
			"Title":     model.Title,
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
