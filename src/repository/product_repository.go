package repository

import (
	"database/sql"

	"github.com/fajrulaulia/arsitektur-bersih/config"
	"github.com/fajrulaulia/arsitektur-bersih/src/model"
)

type ProductRepositoryIface interface {
	Create(code string, name string, price float64) error
	GetByID(id string) (*model.Product, error)
}

type ProductRepositoryStruct struct {
	Config *config.Config
}

func NewProductRepository(c *config.Config) ProductRepositoryIface {
	return &ProductRepositoryStruct{
		Config: c,
	}
}

func (c *ProductRepositoryStruct) Create(code string, name string, price float64) error {
	_, err := c.Config.Db.MySQL().Exec("INSERT INTO products(code, name, price) values (?,?,?)", code, name, price)
	return err
}

func (c *ProductRepositoryStruct) GetByID(id string) (*model.Product, error) {
	var productData model.Product
	var code sql.NullString
	var name sql.NullString
	var price sql.NullFloat64
	var created_at, updated_at sql.NullTime
	row := c.Config.Db.MySQL().QueryRow("SELECT code, name, price, created_at, updated_at FROM products WHERE id = ? or code = ? ", id, id)
	if err := row.Scan(&code, &name, &price, &created_at, &updated_at); err != nil {
		return &productData, err
	}

	productData.Code = code.String
	productData.Name = name.String
	productData.Price = price.Float64
	productData.CreatetAt = created_at.Time
	productData.UpdatedAt = updated_at.Time

	return &productData, nil

}
