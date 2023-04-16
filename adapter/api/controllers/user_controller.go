package controllers

import (
	"net/http"
	"strconv"

	"financial-api/application/dtos"
	"financial-api/application/services"
	"financial-api/domain/entities"
	"financial-api/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// @Summary Get all users
// @Description Get all users
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.User
// @Failure 500 {object} entities.Error
// @Router /api/users [get]
// @Security ApiKeyAuth
func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.userService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// @Summary Create a new user
// @Description Create a new user with the given input data
// @Accept  json
// @Produce  json
// @Param user body entities.User true "User"
// @Success 201 {object} entities.User
// @Failure 400 {object} entities.Error
// @Router /api/users [post]
// @Security ApiKeyAuth
func (c *UserController) Create(ctx *gin.Context) {
	var userDto dtos.UserDTO
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}

	err = c.userService.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// @Summary Update a user
// @Description Update a user with the given input data
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body entities.User true "User"
// @Success 200 {object} entities.User
// @Failure 400 {object} entities.Error
// @Failure 500 {object} entities.Error
// @Router /api/users/{id} [put]
// @Security ApiKeyAuth
func (c *UserController) Update(ctx *gin.Context) {
	var userDto dtos.UserDTO
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		Name:  userDto.Name,
		Email: userDto.Email,
	}

	err = c.userService.Update(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Delete a user
// @Description Delete a user with the given input data
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 204 {object} entities.User
// @Failure 400 {object} entities.Error
// @Failure 500 {object} entities.Error
// @Router /api/users/{id} [delete]
// @Security ApiKeyAuth
func (c *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.userService.Delete(newId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// @Summary Login
// @Description Login with the given input data
// @Accept  json
// @Produce  json
// @Param user body dtos.LoginDTO true "Login"
// @Success 200 {object} entities.JwtToken
// @Failure 401 {object} entities.Error
// @Failure 500 {object} entities.Error
// @Router /api/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var loginInput dtos.LoginDTO

	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.Login(loginInput.Username, loginInput.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(loginInput.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
