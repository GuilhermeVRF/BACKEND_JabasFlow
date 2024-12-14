package controllers

import (
	"jabas-flow/models"
	"jabas-flow/services"
	"jabas-flow/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userService services.UserService
	validator *validator.Validate
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		userService: userService,
		validator: validator.New(),
	}
}

func(userController *UserController) Index (context *gin.Context) {
	users, err := userController.userService.GetUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.NewErrorResponse("error", err.Error()))
	}

	context.JSON(http.StatusOK, users)
}

func (UserController *UserController) Show (context *gin.Context){
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil{
		context.JSON(http.StatusBadRequest, utils.NewErrorResponse("error", err.Error()))	
	}

	var user models.User
	user, err = UserController.userService.GetUser(id)
	if err != nil{
		context.JSON(http.StatusInternalServerError, utils.NewErrorResponse("error", err.Error()))
		return	
	}

	context.JSON(http.StatusOK, user)
}

func (userController *UserController) Store(context *gin.Context) {
	var user models.User	
	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewErrorResponse("error", err.Error()))
		return
	}

	err = userController.validator.Struct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewErrorResponse("error", err.Error()))
		return
	}

	user, err = userController.userService.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.NewErrorResponse("error", err.Error()))
		return
	}

	context.JSON(http.StatusCreated, user)
}

func (userController *UserController) Update(context *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(context.Param("id"))
	
	_, err = userController.userService.GetUser(id)
	if err != nil{
		context.JSON(http.StatusNotFound, utils.NewErrorResponse("error", "Usuário de ID "+ strconv.Itoa(id) + " não encontrado!"))
		return	
	}

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewErrorResponse("error", err.Error()))
		return
	}
	user.ID = id
	err = context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewErrorResponse("error", err.Error()))
	}

	err = userController.validator.Struct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewErrorResponse("error", err.Error()))
		return
	}

	user, err = userController.userService.UpdateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.NewErrorResponse("error", err.Error()))
		return
	}

	context.JSON(http.StatusOK, user)
}