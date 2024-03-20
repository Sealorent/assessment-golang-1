package comment_repo

import "final_project/model"

type ICommentRepository interface {
	Create(comment model.Comment) (model.Comment, error)
}
