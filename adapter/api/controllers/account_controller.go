package controllers

import (
	"financial-api/application/dtos"
	"financial-api/application/services"
	"financial-api/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService *services.AccountService
}

func NewAccountController(accountService *services.AccountService) *AccountController {
	return &AccountController{accountService: accountService}
}

// @Summary Get all accounts
// @Description Get all accounts
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.Account
// @Failure 500 {object} entities.Error
// @Router /api/accounts [get]
// @Security ApiKeyAuth
func (c *AccountController) GetAll(ctx *gin.Context) {
	accounts, err := c.accountService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

// @Summary Create a new account
// @Description Create a new account with the given input data
// @Accept  json
// @Produce  json
// @Param account body entities.Account true "Account"
// @Success 201 {object} entities.Account
// @Failure 400 {object} entities.Error
// @Router /api/accounts [post]
// @Security ApiKeyAuth
func (c *AccountController) Create(ctx *gin.Context) {
	var accountDto dtos.AccountDTO
	err := ctx.ShouldBindJSON(&accountDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := entities.Account{
		OwnerID: accountDto.OwnerID,
		Balance: accountDto.Balance,
	}

	err = c.accountService.Create(&account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, account)
}

// @Summary Update an account
// @Description Update an account with the given input data
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Param account body entities.Account true "Account"
// @Success 200 {object} entities.Account
// @Failure 400 {object} entities.Error
// @Router /api/accounts/{id} [put]
// @Security ApiKeyAuth
func (c *AccountController) Update(ctx *gin.Context) {
	var accountDto dtos.AccountDTO
	err := ctx.ShouldBindJSON(&accountDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := entities.Account{
		OwnerID: accountDto.OwnerID,
		Balance: accountDto.Balance,
	}

	err = c.accountService.Update(&account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// @Summary Delete an account
// @Description Delete an account with the given input data
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 204
// @Failure 400 {object} entities.Error
// @Failure 500 {object} entities.Error
// @Router /api/accounts/{id} [delete]
// @Security ApiKeyAuth
func (c *AccountController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.accountService.Delete(newId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
