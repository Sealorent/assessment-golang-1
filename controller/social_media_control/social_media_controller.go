package social_media_control

import (
	"final_project/common"
	"final_project/model"
	"final_project/repository/social_media_repo"
	"final_project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	socialMediaRepository social_media_repo.ISocialMediaRepository
}

func NewSocialMediaController(socialMediaRepository social_media_repo.ISocialMediaRepository) *SocialMediaController {
	return &SocialMediaController{
		socialMediaRepository: socialMediaRepository,
	}
}

func (smc *SocialMediaController) Create(ctx *gin.Context) {
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

	// map input to user struct
	var newSocialMedia model.SocialMedia

	// bind the input to the user struct
	errors := ctx.ShouldBindJSON(&newSocialMedia)
	if errors != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: err.Error(),
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
		return
	}

	newSocialMedia.UserId = userId
	createdSocialMedia, err := smc.socialMediaRepository.Create(newSocialMedia)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to create social media",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
		return
	}

	response := model.SocialMediaResultCreated{
		ID:             createdSocialMedia.Id,
		Name:           createdSocialMedia.Name,
		SocialMediaUrl: createdSocialMedia.SocialMediaUrl,
		UserId:         createdSocialMedia.UserId,
		CreatedAt:      createdSocialMedia.CreatedAt.String(),
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Success to create social media",
		Data:    response,
	}

	ctx.JSON(http.StatusCreated, r)

}

func (smc *SocialMediaController) FindAll(ctx *gin.Context) {
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

	socialMedia, err := smc.socialMediaRepository.FindAll()
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to find social media",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(400, r)
		return
	}

	var socialMediaResults []model.SocialMediaResult

	for i := 0; i < len(socialMedia); i++ {
		socialMediaResult := model.SocialMediaResult{
			ID:             socialMedia[i].Id,
			Name:           socialMedia[i].Name,
			SocialMediaUrl: socialMedia[i].SocialMediaUrl,
			UserId:         socialMedia[i].UserId,
			CreatedAt:      socialMedia[i].CreatedAt.String(),
			UpdatedAt:      socialMedia[i].UpdatedAt.String(),
			User: model.UserReferSocialMedia{
				ID:       socialMedia[i].User.ID,
				Username: socialMedia[i].User.Username,
			},
		}

		// // Append the populated socialMediaResult to the socialMediaResults slice
		socialMediaResults = append(socialMediaResults, socialMediaResult)
	}

	var length int = len(socialMediaResults)
	var r common.ResponseSocialMedia
	if length > 0 {
		r = common.ResponseSocialMedia{
			Success: true,
			Message: "Success",
			Data: common.SocialMediaDTO{
				SocialMedia: socialMediaResults,
			},
		}
	} else {
		r = common.ResponseSocialMedia{
			Success: false,
			Message: "No social media found",
			Data:    common.SocialMediaDTO{}, // You might want to provide an empty DTO here or handle this case appropriately.
		}
	}

	ctx.JSON(http.StatusOK, r)
}

func (smc *SocialMediaController) FindOne(ctx *gin.Context) {
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

	socialMediaId := ctx.Param("socialMediaId")
	socialMedia, err := smc.socialMediaRepository.FindOne(socialMediaId)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to find social media",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	response := model.SocialMediaResult{
		ID:             socialMedia.Id,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		CreatedAt:      socialMedia.CreatedAt.String(),
		UpdatedAt:      socialMedia.UpdatedAt.String(),
		User: model.UserReferSocialMedia{
			ID:       socialMedia.User.ID,
			Username: socialMedia.User.Username,
		},
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Success",
		Data:    response,
	}

	ctx.JSON(http.StatusOK, r)

}

func (smc *SocialMediaController) UpdateOne(ctx *gin.Context) {
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

	// map input to user struct
	var request model.SocialMedia
	if err := ctx.ShouldBindJSON(&request); err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to update social media",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	socialMediaId := ctx.Param("socialMediaId")
	updatedSocialMedia, err := smc.socialMediaRepository.UpdateOne(request, socialMediaId, userId)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to update social media",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	response := model.SocialMediaResultUpdated{
		ID:             updatedSocialMedia.Id,
		Name:           updatedSocialMedia.Name,
		SocialMediaUrl: updatedSocialMedia.SocialMediaUrl,
		UserId:         updatedSocialMedia.UserId,
		UpdatedAt:      updatedSocialMedia.UpdatedAt.String(),
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Success to update social media",
		Data:    response,
	}

	ctx.JSON(http.StatusOK, r)

}

func (smc *SocialMediaController) Delete(ctx *gin.Context) {
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

	socialMediaId := ctx.Param("socialMediaId")
	err = smc.socialMediaRepository.Delete(socialMediaId, userId)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to delete social media",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	var r common.Response = common.Response{
		Success: true,
		Message: "Your social media has been successfully deleted",
	}

	ctx.JSON(http.StatusOK, r)

}
