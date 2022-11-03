package main

import (
	"github.com/fajrulaulia/arsitektur-bersih/config"
	"github.com/fajrulaulia/arsitektur-bersih/src/delivery"
	"github.com/fajrulaulia/arsitektur-bersih/src/repository"
	"github.com/fajrulaulia/arsitektur-bersih/src/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	cfg := config.InitConfig()
	productRepos := repository.NewProductRepository(cfg)
	productUsecase := usecase.NewProductUsecase(productRepos)
	deliverProduct := delivery.NewProductDelivery(productUsecase)

	e := echo.New()

	deliverProduct.Apply(e)

	e.Logger.Fatal(e.Start(":1323"))

}
