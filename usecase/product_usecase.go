package usecase

import (
	"github.com/natanpms/remember-api/model"
	"github.com/natanpms/remember-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	return pu.repository.CreateProduct(product)
}

func (pu *ProductUseCase) GetProductByID(id int) (model.Product, error) {
	return pu.repository.GetProductByID(id)
}

func (pu *ProductUseCase) DeleteProduct(id int) error {
	return pu.repository.DeleteProduct(id)
}
