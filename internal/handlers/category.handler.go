package handlers

import (
	"backEnd_Coffeshop/internal/models"
	"backEnd_Coffeshop/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerCategory struct {
	*repositories.RepoCategory
}

func NewCategory(r *repositories.RepoCategory) *HandlerCategory {
	return &HandlerCategory{r}
}

func (h *HandlerCategory) GetCategories(ctx *gin.Context) {
	categories, err := h.GetCategory()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (h *HandlerCategory) CreateCategory(ctx *gin.Context) {
	var category models.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.RepoCategory.CreateCategory(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": response})

}

func (h *HandlerCategory) DeleteCategory(ctx *gin.Context) {
	// Get the product ID from the URL path
	idCategorys := ctx.Param("id")

	if idCategorys == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Category ID not provided"})
		return
	}

	response, err := h.RepoCategory.DeleteCategory(idCategorys)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": response})
}
