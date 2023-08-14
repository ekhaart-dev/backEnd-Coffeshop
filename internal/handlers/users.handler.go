package handlers

import (
	"backEnd_Coffeshop/config"
	"backEnd_Coffeshop/internal/models"
	"backEnd_Coffeshop/internal/repositories"
	"backEnd_Coffeshop/pkg"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUsers struct {
	*repositories.RepoUsers
}

func NewUsers(r *repositories.RepoUsers) *HandlerUsers {
	return &HandlerUsers{r}
}

func (h *HandlerUsers) PostData(ctx *gin.Context) {
	var ers error
	data := models.Users{
		Role: "user",
	}

	if ers = ctx.ShouldBind(&data); ers != nil {
		ctx.AbortWithError(http.StatusBadRequest, ers)
		return
	}

	_, ers = govalidator.ValidateStruct(&data)
	if ers != nil {
		pkg.NewRes(401, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	data.Password, ers = pkg.HashPassword(data.Password)
	if ers != nil {
		pkg.NewRes(401, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	respone, ers := h.CreateUsers(&data)
	if ers != nil {
		ctx.AbortWithError(http.StatusBadRequest, ers)
		return
	}

	pkg.NewRes(200, respone).Send(ctx)

}

func (h *HandlerUsers) FetchAll(ctx *gin.Context) {
	data, err := h.GetAllUser()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, data).Send(ctx)
}

func (h *HandlerUsers) UpdateUsers(ctx *gin.Context) {
	var users models.Users
	users.Id_user = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.RepoUsers.UpdateUsers(&users)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"message":     response,
	})
}

func (h *HandlerUsers) DeleteUser(ctx *gin.Context) {
	// Get the product ID from the URL path
	idUser := ctx.Param("id")

	if idUser == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Users ID not provided"})
		return
	}

	response, err := h.RepoUsers.DeleteUser(idUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete users"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": response})
}
