package api

import (
	"net/http"
	"restsample/db/models"
	"restsample/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepo repository.UserRepo
}

type IUserController interface {
	GetUser() gin.HandlerFunc
	GetAllUsers() gin.HandlerFunc
}

type userInput struct {
	Ids []int `json:"ids" binding:"required"`
}

func (controller *userController) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqId := ctx.Param("id")
		id, err := strconv.Atoi(reqId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "enter valid id",
			})
			return
		}
		var userQuery models.UserQuery
		userQuery.Id = &id
		user, err := controller.userRepo.FindOne(userQuery)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong",
			})
			return
		}
		if user == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data": "no user found ",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": user,
		})
	}
}
func (controller *userController) GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input userInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ids := make(map[int]bool)
		for _, x := range input.Ids {
			ids[x] = true
		}
		var query models.UserQuery
		query.IdsMap = ids
		users, err := controller.userRepo.FindAll(query)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong",
			})
			return
		}
		if len(users) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"data": "no users found ",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	}
}

func NewUserController(userRepo repository.UserRepo) IUserController {
	return &userController{
		userRepo: userRepo,
	}
}
