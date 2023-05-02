package services

import (
	"GoCommunityAPI/models"
	"GoCommunityAPI/repositories"
	"errors"
)

func FetchCommentPageByArticleId(articleId int, page int, pageSize int) ([]models.CommentModel, error) {
	comments, err := repositories.FetchCommentPageByArticleId(articleId, page, pageSize)
	return comments, err
}

func UploadComment(comment models.CommentModel) error {
	_, err := repositories.FindOneUserById(comment.UserId)
	if err != nil {
		return errors.New("You must provide a valid user id")
	}
	_, err = repositories.FindOneArticle(comment.ArticleId)
	if err != nil {
		return errors.New("You must provide a valid article id")
	}
	err = repositories.CreateComment(comment)
	return err
}

func DeleteComment(commentId int) error {
	_, err := repositories.FindOneComment(commentId)
	if err != nil {
		return errors.New("You must provide a valid comment id")
	}
	err = repositories.DeleteComment(commentId)
	return err
}
