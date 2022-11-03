package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/fajrulaulia/arsitektur-bersih/src/model"
	"github.com/fajrulaulia/arsitektur-bersih/src/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFailOKProduct(t *testing.T) {
	timeData, _ := time.Parse(time.RFC1123, "2022-11-03T07:37:45Z")
	productRepoMock := &mocks.ProductRepositoryIface{}
	productRepoMock.On("GetByID", mock.Anything).Return(&model.Product{
		Code:      "1234",
		Name:      "hola",
		Price:     3000,
		UpdatedAt: timeData,
	}, nil)
	productUsecase := NewProductUsecase(productRepoMock)

	res, err := productUsecase.AmbilBarangBerdasarkanID("1")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestFailGetProduct(t *testing.T) {
	productRepoMock := &mocks.ProductRepositoryIface{}
	productRepoMock.On("GetByID", mock.Anything).Return(nil, errors.New("Data not found"))
	productUsecase := NewProductUsecase(productRepoMock)

	res, err := productUsecase.AmbilBarangBerdasarkanID("1")
	assert.Nil(t, res)
	assert.NotNil(t, err)
}
