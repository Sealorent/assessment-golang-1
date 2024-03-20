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

	sub, err := utils.CheckTokenJWTAndReturnSub(ctx)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized",
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
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
		return
	}

	newPhoto.UserID = sub.(string)
	createdPhoto, err := uc.photoRepository.Create(newPhoto)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to create photo",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
		return
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
		UserId:    createdPhoto.UserID,
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
		return
	}
	var photoResults []model.PhotoResult

	for i := 0; i < len(photos); i++ {
		photoResult := model.PhotoResult{
			ID:        photos[i].Id,
			Title:     photos[i].Title,
			Caption:   photos[i].Caption,
			PhotoUrl:  photos[i].PhotoUrl,
			UserID:    photos[i].UserID,
			CreatedAt: photos[i].CreatedAt.String(),
			UpdateAt:  photos[i].UpdatedAt.String(),
			User: model.UserRefer{
				Username: photos[i].User.Username,
				Email:    photos[i].User.Email,
			},
		}

		// Append the populated photoResult to the photoResults slice
		photoResults = append(photoResults, photoResult)
	}

	var r common.Response = common.CreateResponse(true, "Success", photoResults, "")
	ctx.JSON(http.StatusOK, r)
}

func (uc *PhotoController) UpdateOne(ctx *gin.Context) {
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

	var request model.Photo
	errors := ctx.ShouldBindJSON(&request)
	if errors != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Process Failed",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}
	var photoId = ctx.Query("photoId")
	var userId = sub.(string)

	updatePhoto, err := uc.photoRepository.UpdateOne(request, photoId, userId)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to Update",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	type UpdatedPhotoResponse struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Caption   string `json:"caption"`
		PhotoUrl  string `json:"photo_url"`
		UserId    string `json:"user_id"`
		UpdatedAt string `json:"created_at"`
	}

	response := UpdatedPhotoResponse{
		ID:        updatePhoto.Id,
		Title:     updatePhoto.Title,
		Caption:   updatePhoto.Caption,
		PhotoUrl:  updatePhoto.PhotoUrl,
		UserId:    updatePhoto.UserID,
		UpdatedAt: updatePhoto.UpdatedAt.String(),
	}

	var r common.Response = common.CreateResponse(true, "Photo updated successfully", response, "")

	ctx.JSON(http.StatusOK, r)

}

func (uc *PhotoController) Delete(ctx *gin.Context) {
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

	var photoId = ctx.Query("photoId")
	var userId = sub.(string)

	err = uc.photoRepository.Delete(photoId, userId)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to Delete",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	var r common.Response = common.CreateResponse(true, "Photo deleted successfully", nil, "")

	ctx.JSON(http.StatusOK, r)

}
