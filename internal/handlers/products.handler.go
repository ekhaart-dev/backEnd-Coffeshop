package handlers

import (
	"backEnd_Coffeshop/internal/models"
	"backEnd_Coffeshop/internal/repositories"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	*repositories.RepoProduct
}

func NewProductHandler(productRepo *repositories.RepoProduct) *ProductHandler {
	return &ProductHandler{productRepo}
}

func (h *ProductHandler) GetProducts(ctx *gin.Context) {
	var meta_product models.Meta_Products
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	search := ctx.Query("search")
	orderby := ctx.Query("order_by")
	var offset int = 0
	var page_int, _ = strconv.Atoi(page)
	var limit_int, _ = strconv.Atoi(limit)
	if limit == "" {
		limit_int = 5
	}
	if page == "" {
		page_int = 1
	}
	if page_int > 0 {
		offset = (page_int - 1) * limit_int
	} else {
		offset = 0
	}

	count_data := h.Get_Count_Data(search)

	if count_data <= 0 {
		meta_product.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_product.Next = ""
		} else {
			meta_product.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_product.Prev = ""
	} else {
		meta_product.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_product.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_product.Last_page = ""
	}

	if count_data != 0 {
		meta_product.Total_data = strconv.Itoa(count_data)
	} else {
		meta_product.Last_page = ""
	}

	// Panggil fungsi GetBy dengan parameter yang sesuai
	products, err := h.GetBy(offset, limit_int, search, orderby)
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
		"data":        products, // Menggunakan response yang diperoleh dari GetBy
		"meta":        meta_product,
	})
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var productset models.Products

	if err := ctx.ShouldBind(&productset); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Requestt",
			"message":     err.Error(),
		})
		return
	}

	productset.Banner_product = ctx.MustGet("banner_product").(string)
	response, err := h.CreateProductData(&productset)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
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

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var product models.Products
	product.Id_product = ctx.Param("id")
	product.Banner_product = ctx.MustGet("banner_product").(string)

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.RepoProduct.UpdateProduct(&product)
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

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	// Get the product ID from the URL path
	idProduct := ctx.Param("id")

	if idProduct == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product ID not provided"})
		return
	}

	response, err := h.RepoProduct.DeleteProduct(idProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": response})
}
