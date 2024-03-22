package comment_control

import (
	"final_project/common"
	"final_project/model"
	"final_project/repository/comment_repo"
	"final_project/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentRepository comment_repo.ICommentRepository
}

func NewCommentController(commentRepository comment_repo.ICommentRepository) *commentController {
	return &commentController{
		commentRepository: commentRepository,
	}
}

// CreateComment is a function to create a new comment
func (cc *commentController) CreateComment(ctx *gin.Context) {

	userId, err := utils.CheckTokenJWTAndReturnSub(ctx)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	if userId == "" {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   "Unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	var newComment model.Comment
	err = ctx.ShouldBindJSON(&newComment)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: err.Error(),
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	if err := newComment.Validate(); err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Please recheck your input : " + err.Error(),
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	newComment.UserId = userId
	createdComment, err := cc.commentRepository.Create(newComment)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to create comment",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
		return
	}

	response := model.CommentResultCreate{
		ID:        createdComment.Id,
		Message:   createdComment.Message,
		PhotoId:   createdComment.PhotoId,
		UserID:    createdComment.UserId,
		CreatedAt: createdComment.CreatedAt.String(),
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Comment created",
		Data:    response,
	}

	ctx.JSON(http.StatusCreated, r)
}

func (cc *commentController) GetAll(ctx *gin.Context) {
	sub, err := utils.CheckTokenJWTAndReturnSub(ctx)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	if sub == "" {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	comments, err := cc.commentRepository.GetAll()
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to get comments",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	var commentResults []model.CommentResult

	for i := 0; i < len(comments); i++ {
		commentResult := model.CommentResult{
			ID:        comments[i].Id,
			Message:   comments[i].Message,
			PhotoId:   comments[i].PhotoId,
			UserID:    comments[i].UserId,
			CreatedAt: comments[i].CreatedAt.String(),
			UpdatedAt: comments[i].UpdatedAt.String(),
			User: model.UserReferComment{
				ID:       comments[i].User.ID,
				Username: comments[i].User.Username,
				Email:    comments[i].User.Email,
			},
			Photo: model.PhotoReferComment{
				ID:       comments[i].Photo.Id,
				Title:    comments[i].Photo.Title,
				Caption:  comments[i].Photo.Caption,
				PhotoUrl: comments[i].Photo.PhotoUrl,
				UserID:   comments[i].Photo.UserID,
			},
		}

		commentResults = append(commentResults, commentResult)
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Success",
		Data:    commentResults,
	}

	ctx.JSON(http.StatusOK, r)

}

func (cc *commentController) GetOne(ctx *gin.Context) {
	userId, err := utils.CheckTokenJWTAndReturnSub(ctx)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	if userId == "" {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   "Unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	commentId := ctx.Param("commentId")

	comment, err := cc.commentRepository.GetOne(commentId)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to get comment",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	response := model.CommentResult{
		ID:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserID:    comment.UserId,
		CreatedAt: comment.CreatedAt.String(),
		UpdatedAt: comment.UpdatedAt.String(),
		User: model.UserReferComment{
			ID:       comment.User.ID,
			Username: comment.User.Username,
			Email:    comment.User.Email,
		},
		Photo: model.PhotoReferComment{
			ID:       comment.Photo.Id,
			Title:    comment.Photo.Title,
			Caption:  comment.Photo.Caption,
			PhotoUrl: comment.Photo.PhotoUrl,
			UserID:   comment.Photo.UserID,
		},
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Success",
		Data:    response,
	}

	ctx.JSON(http.StatusOK, r)
}

func (cc *commentController) Update(ctx *gin.Context) {
	userId, err := utils.CheckTokenJWTAndReturnSub(ctx)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	var comment model.Comment
	err = ctx.ShouldBindJSON(&comment)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: err.Error(),
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	commentId := ctx.Param("commentId")
	updatedComment, err := cc.commentRepository.Update(comment, commentId, userId)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to update comment",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
		return
	}

	fmt.Println("updatedComment", updatedComment)

	response := model.CommentResultUpdate{
		ID:        updatedComment.Id,
		Title:     updatedComment.Photo.Title,
		Caption:   updatedComment.Photo.Caption,
		PhotoUrl:  updatedComment.Photo.PhotoUrl,
		UserID:    updatedComment.Photo.UserID,
		UpdatedAt: updatedComment.UpdatedAt.String(),
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Comment updated with message: " + updatedComment.Message,
		Data:    response,
	}

	ctx.JSON(http.StatusOK, r)
}

func (cc *commentController) Delete(ctx *gin.Context) {
	userId, err := utils.CheckTokenJWTAndReturnSub(ctx)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	commentId := ctx.Param("commentId")

	err = cc.commentRepository.Delete(commentId, userId)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to delete comment",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
		return
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Your comment has been successfully deleted",
	}

	ctx.JSON(http.StatusOK, r)
}
