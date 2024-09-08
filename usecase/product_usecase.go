package usecase

import "go-api/model"

type ProductUsecase struct {
	//Repository
}

func NewProductUseCase() ProductUsecase {
	return ProductUsecase{}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
