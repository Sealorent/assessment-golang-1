// make package controller and folder user and file UserController.go
package user_control

import (
	"final_project/common"
	"final_project/model"
	"final_project/repository/user_repo"
	"final_project/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository user_repo.IUserRepository
}

func NewUserController(userRepository user_repo.IUserRepository) *UserController {
	return &UserController{
		userRepository: userRepository,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {

	// map input to user struct
	var newUser model.User

	// bind the input to the user struct
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: err.Error(),
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	// hash the password
	hashedPassword, err := utils.Hash([]byte(newUser.Password))
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Error hashing password",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	newUser.Password = string(hashedPassword)
	registeredUser, err := uc.userRepository.Register(newUser)
	fmt.Println(err)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to Register",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	var registeredUserResponse = gin.H{
		"id":       registeredUser.ID,
		"username": registeredUser.Username,
		"email":    registeredUser.Email,
		"age":      registeredUser.Age,
	}

	var response common.Response = common.CreateResponse(true, "User registered successfully", registeredUserResponse, "")

	ctx.JSON(http.StatusCreated, response)
}

func (uc *UserController) Login(ctx *gin.Context) {
	var requestedUser model.User

	err := ctx.ShouldBindJSON(&requestedUser)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Process Failed",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	user, err := uc.userRepository.UserByEmail(requestedUser.Email)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Data Not Found",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, r)
		return
	}

	if !utils.HashMatched([]byte(user.Password), []byte(requestedUser.Password)) {
		var r common.Response = common.Response{
			Success: false,
			Message: "Unauthorized Please Check your Password",
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	token, err := utils.GenerateJWTToken(user.ID, user.Email, user.Username)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to generate token",
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	var response = gin.H{
		"token": token,
	}

	var r common.Response = common.CreateResponse(true, "Login Success", response, "")

	ctx.JSON(http.StatusOK, r)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

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

	var request model.User
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
	var id = ctx.Param("userId")
	if id != sub.(string) {
		var r common.Response = common.Response{
			Success: false,
			Message: "Cannot Update other user",
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	updateUser, err := uc.userRepository.UpdateUser(request, id)

	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to Update",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	var response = gin.H{
		"id":         updateUser.ID,
		"username":   updateUser.Username,
		"email":      updateUser.Email,
		"age":        updateUser.Age,
		"updated_at": updateUser.UpdatedAt,
	}

	var r common.Response = common.CreateResponse(true, "User updated successfully", response, "")

	ctx.JSON(http.StatusOK, r)

}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

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

	err = uc.userRepository.DeleteUser(sub.(string))
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Failed to Delete",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	var r common.Response = common.CreateResponse(true, "Your account has been successfully deleted", nil, "")

	ctx.JSON(http.StatusOK, r)
}
