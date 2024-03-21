package comment_repo

import (
	"errors"
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

	// Validate the input
	if comment.PhotoId == 0 {
		return model.Comment{}, errors.New("photo_id is required")
	}

	err := comment.Validate()
	if err != nil {
		return model.Comment{}, err
	}

	tx := ur.db.Create(&comment)
	if tx.Error != nil {
		return model.Comment{}, tx.Error
	}

	// Return the created comment
	// find comment by id
	var createdComment model.Comment
	tx = ur.db.Where("id = ?", comment.Id).Preload("User").Preload("Photo").First(&createdComment)
	if tx.Error != nil {
		return model.Comment{}, tx.Error
	}

	return createdComment, nil

}

func (ur *commentRepository) Update(comment model.Comment, commentId string, userId string) (model.Comment, error) {

	// Validate the input
	err := comment.Validate()
	if err != nil {
		return model.Comment{}, err
	}

	// Make begin transaction
	tx := ur.db.Begin()

	// 1. Find comment by id
	var existingComment model.Comment
	if err := tx.Where("id = ? AND user_id = ? AND status = true", commentId, userId).First(&existingComment).Error; err != nil {
		tx.Rollback()
		return model.Comment{}, errors.New("comment not found or you don't have permission to update this comment")
	}

	// 2. Update comment
	tx = tx.Model(&existingComment).Updates(comment)
	if tx.Error != nil {
		tx.Rollback()
		return model.Comment{}, errors.New("failed to update comment")
	}

	// 3. Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return model.Comment{}, errors.New("failed to commit transaction")
	}

	// 4. Return updated comment
	var updatedComment model.Comment
	tx = ur.db.Where("id = ?", commentId).Preload("User").Preload("Photo").First(&updatedComment)
	if tx.Error != nil {
		return model.Comment{}, errors.New("failed to get updated comment")
	}

	return updatedComment, nil

}

func (cc *commentRepository) GetAll() ([]model.Comment, error) {
	var comments []model.Comment
	tx := cc.db.Where("status = ?", true).Preload("User").Preload("Photo").Find(&comments)
	if tx.Error != nil {
		return []model.Comment{}, tx.Error
	}
	return comments, nil
}

func (cc *commentRepository) GetOne(commentId string) (model.Comment, error) {

	var comment model.Comment

	tx := cc.db.Where("id = ?", commentId).Preload("User").Preload("Photo").Find(&comment)
	if tx.Error != nil {
		return model.Comment{}, tx.Error
	}

	if !comment.Status {
		return model.Comment{}, errors.New("comment not found")
	}

	return comment, nil
}

func (cc *commentRepository) Delete(commentId string, userId string) error {

	// Make begin transaction
	tx := cc.db.Begin()

	// 1. Find comment by id and user_id
	var existingComment model.Comment
	if err := tx.Where("id = ? AND user_id = ? AND status = true", commentId, userId).First(&existingComment).Error; err != nil {
		tx.Rollback()
		return errors.New("comment not found or you don't have permission to delete this comment")
	}

	// 2. Soft delete comment
	tx = tx.Model(&existingComment).Update("status", false)
	if tx.Error != nil {
		tx.Rollback()
		return errors.New("failed to delete comment")
	}

	// 3. Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}
