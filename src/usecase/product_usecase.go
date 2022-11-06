package usecase

import (
	"github.com/fajrulaulia/arsitektur-bersih/src/repository"
	"github.com/fajrulaulia/arsitektur-bersih/src/usecase/products"
)

type ProductUsecaseIface interface {
	InsertBarangBaru(code string, name string, price float64) error
	AmbilBarangBerdasarkanID(id string) (*products.ProductResponse, error)
}

type ProductUsecaseStruct struct {
	ProductRepo repository.ProductRepositoryIface
}

func NewProductUsecase(product repository.ProductRepositoryIface) ProductUsecaseIface {
	return &ProductUsecaseStruct{
		ProductRepo: product,
	}
}

func (c *ProductUsecaseStruct) InsertBarangBaru(code string, name string, price float64) error {
	err := c.ProductRepo.Create(code, name, price)
	if err != nil {
		return err
	}
	return err
}

func (c *ProductUsecaseStruct) AmbilBarangBerdasarkanID(id string) (*products.ProductResponse, error) {
	var product products.ProductResponse
	res, err := c.ProductRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	product.Code = res.Code
	product.Name = res.Name
	product.Price = res.Price
	product.CreatedAt = res.CreatetAt
	product.UpdatedAt = res.UpdatedAt

	return &product, nil
}
