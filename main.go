package main

import (
	"gofr.dev/cmd/gofr/migration"

	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
	"gofr.dev/pkg/gofr"

	"products/handler/products"
	"products/handler/variants"
	"products/migrations"

	productsService "products/service/products"
	variantService "products/service/variants"
	productsStore "products/store/products"
	variantsStore "products/store/variants"
)

func main() {
	app := gofr.New()

	db := dbmigration.NewGorm(app.GORM())

	err := migration.Migrate("products", db, migrations.All(), "UP", app.Logger)
	if err != nil {
		app.Logger.Errorf("Migration failed with error: ", err)
	}

	variantStore := variantsStore.New()
	productStore := productsStore.New(variantStore)

	productSvc := productsService.New(productStore, variantStore)
	variantSvc := variantService.New(productStore, variantStore)

	productHandler := products.New(productSvc)
	variantHandler := variants.New(variantSvc)

	app.POST("/product", productHandler.Create)
	app.GET("/product/{id}", productHandler.GetByID)
	app.GET("/products", productHandler.GetAll)

	app.GET("/products/{pid}/variant/{id}", variantHandler.GetByID)
	app.POST("/products/{pid}/variant", variantHandler.Create)

	app.Start()
}
