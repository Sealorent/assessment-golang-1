package comment_repo

import "final_project/model"

type ICommentRepository interface {
	Create(comment model.Comment) (model.Comment, error)
	Update(comment model.Comment, commentId string, userId string) (model.Comment, error)
	GetAll() ([]model.Comment, error)
	GetOne(commentId string) (model.Comment, error)
	Delete(commentId string, userId string) error
}
