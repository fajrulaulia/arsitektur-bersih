package delivery

import (
	"log"
	"net/http"

	"github.com/fajrulaulia/arsitektur-bersih/src/usecase"
	"github.com/fajrulaulia/arsitektur-bersih/src/usecase/products"
	"github.com/labstack/echo/v4"
)

type ProductDelivery struct {
	Product usecase.ProductUsecaseIface
}

func NewProductDelivery(c usecase.ProductUsecaseIface) ProductDelivery {
	return ProductDelivery{
		Product: c,
	}
}

func (d *ProductDelivery) Apply(e *echo.Echo) *echo.Echo {
	e.POST("/product", d.Create)
	e.GET("/product/:id", d.Get)

	return e
}

func (d *ProductDelivery) Create(c echo.Context) error {
	data := new(products.ProductRequest)

	if err := c.Bind(data); err != nil {
		return err
	}

	if err := d.Product.InsertBarangBaru(data.Code, data.Name, data.Price); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, data)

}

func (d *ProductDelivery) Get(c echo.Context) error {

	res, err := d.Product.AmbilBarangBerdasarkanID(c.Param("id"))
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, res)
}
