package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/natanpms/remember-api/handlers"
	"github.com/natanpms/remember-api/model"
	"github.com/natanpms/remember-api/usecase"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

// GetProducts
// @Summary Get all products
// @Description Get a list of all products
// @Tags products
// @Produce json
// @Success 200 {object} handlers.APIResponse
// @Failure 500 {object} handlers.APIResponse
// @Router /api/v1/products [get]
func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()

	if err != nil {
		handlers.Error(ctx, 500, "Erro ao buscar produtos", err.Error())
		return
	}

	handlers.Success(ctx, 200, "", products)
}

// CreateProduct
// @Summary Create a product
// @Description Create a new product with the given payload
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product payload"
// @Success 201 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Failure 500 {object} handlers.APIResponse
// @Router /api/v1/products [post]
func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		handlers.InvalidRequest(ctx, "Erro ao ler dados do produto")
		return
	}

	productInserted, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		handlers.Error(ctx, 500, "Erro ao criar produto", err.Error())
		return
	}

	handlers.Success(ctx, 201, "Produto criado com sucesso", productInserted)
}

// GetProductByID
// @Summary Get a product by ID
// @Description Get a single product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} handlers.APIResponse
// @Failure 404 {object} handlers.APIResponse
// @Failure 500 {object} handlers.APIResponse
// @Router /api/v1/products/{id} [get]
func (p *productController) GetProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handlers.Error(ctx, 500, "ID inválido", err.Error())
		return
	}

	product, err := p.productUseCase.GetProductByID(id)

	if err != nil {
		handlers.NotFound(ctx, err.Error())
		return
	}

	handlers.Success(ctx, 200, "", product)
}

// DeleteProduct
// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} handlers.APIResponse
// @Failure 404 {object} handlers.APIResponse
// @Failure 500 {object} handlers.APIResponse
// @Router /api/v1/products/{id} [delete]
func (p *productController) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handlers.Error(ctx, 500, "ID inválido", err.Error())
		return
	}

	err = p.productUseCase.DeleteProduct(id)

	if err != nil {
		handlers.NotFound(ctx, err.Error())
		return
	}

	handlers.Success(ctx, 200, "Produto excluído com sucesso", nil)
}
