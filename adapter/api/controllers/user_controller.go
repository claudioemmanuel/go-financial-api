package controllers

import (
	"net/http"

	"financial-api/application/dtos"
	"financial-api/application/services"
	"financial-api/domain/entities"

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
func (c *UserController) Create(ctx *gin.Context) {
	var userDto dtos.UserDTO
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
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
func (c *UserController) Update(ctx *gin.Context) {
	var userDto dtos.UserDTO
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
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
// @Success 200 {object} entities.User
// @Failure 400 {object} entities.Error
// @Failure 500 {object} entities.Error
// @Router /api/users/{id} [delete]
func (c *UserController) Delete(ctx *gin.Context) {
	var userDto dtos.UserDTO
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
	}

	err = c.userService.Delete(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
