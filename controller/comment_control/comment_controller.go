package comment_control

import (
	"final_project/common"
	"final_project/model"
	"final_project/repository/comment_repo"
	"final_project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	photoRepository comment_repo.ICommentRepository
}

func NewCommentController(photoRepository comment_repo.ICommentRepository) *commentController {
	return &commentController{
		photoRepository: photoRepository,
	}
}

// CreateComment is a function to create a new comment
func (cc *commentController) CreateComment(ctx *gin.Context) {

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

	newComment.UserId = sub
	createdComment, err := cc.photoRepository.Create(newComment)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to create comment",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
		return
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Comment created",
		Data:    createdComment,
	}

	ctx.JSON(http.StatusCreated, r)

}
