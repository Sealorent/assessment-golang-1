package photo_control

import (
	"final_project/common"
	"final_project/model"
	"final_project/repository/photo_repo"
	"final_project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoRepository photo_repo.IPhotoRepository
}

func NewPhotoController(photoRepository photo_repo.IPhotoRepository) *PhotoController {
	return &PhotoController{
		photoRepository: photoRepository,
	}
}

func (uc *PhotoController) Create(ctx *gin.Context) {

	claims, exist := ctx.Get("claims")
	if !exist {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	sub, err := utils.GetSubFromClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.CreateResponse(false, "Unauthorized", nil, "unauthorized"))
		return
	}

	// map input to user struct
	var newPhoto model.Photo

	// bind the input to the user struct
	errors := ctx.ShouldBindJSON(&newPhoto)
	if errors != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: err.Error(),
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
	}

	newPhoto.UserId = sub.(string)
	createdPhoto, err := uc.photoRepository.Create(newPhoto)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to create photo",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
	}

	type CreatedPhotoResponse struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Caption   string `json:"caption"`
		PhotoUrl  string `json:"photo_url"`
		UserId    string `json:"user_id"`
		CreatedAt string `json:"created_at"`
	}

	response := CreatedPhotoResponse{
		ID:        createdPhoto.Id,
		Title:     createdPhoto.Title,
		Caption:   createdPhoto.Caption,
		PhotoUrl:  createdPhoto.PhotoUrl,
		UserId:    createdPhoto.UserId,
		CreatedAt: createdPhoto.CreatedAt.String(),
	}

	var r common.Response = common.CreateResponse(true, "Created successfully", response, "")

	ctx.JSON(http.StatusCreated, r)
}

func (uc *PhotoController) FindAll(ctx *gin.Context) {
	photos, err := uc.photoRepository.FindAll()
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to get photos",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
	}
	var r common.Response = common.CreateResponse(true, "Success", photos, "")
	ctx.JSON(http.StatusOK, r)
}
