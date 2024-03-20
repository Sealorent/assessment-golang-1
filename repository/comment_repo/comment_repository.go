package comment_repo

import (
	"final_project/model"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (ur *commentRepository) Create(comment model.Comment) (model.Comment, error) {

	err := comment.Validate()
	if err != nil {
		return model.Comment{}, err
	}
	tx := ur.db.Begin()

	tx.Create(&comment)
	if tx.Error != nil {
		tx.Rollback()
		return model.Comment{}, tx.Error
	}

	return comment, nil

}
