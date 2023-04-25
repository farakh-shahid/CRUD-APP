package controllers

import (
	"net/http"

	"github.com/farakh-shahid/CRUD-APP/models"
	"github.com/farakh-shahid/CRUD-APP/services"
	"github.com/farakh-shahid/CRUD-APP/utils/constants"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func NewUserController(userservice services.UserServiceInterface) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{constants.MessageKey: err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{constants.MessageKey: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{constants.MessageKey: constants.CreateUser})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	var username string = ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{constants.MessageKey: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{constants.MessageKey: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{constants.MessageKey: err.Error()})
		return
	}

	userID := ctx.Param("id")

	err := uc.UserService.UpdateUser(userID, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{constants.MessageKey: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{constants.MessageKey: constants.UpdateUser})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.UserService.DeleteUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{constants.MessageKey: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{constants.MessageKey: constants.DeleteUser})
}
