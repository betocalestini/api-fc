package main

import (
	"fmt"
	"net/http"

	"github.com/betocalestini/api-fc/configs"
	"github.com/betocalestini/api-fc/internal/entity"
	"github.com/betocalestini/api-fc/internal/infra/database"
	"github.com/betocalestini/api-fc/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/user", userHandler.CreateUser)

	fmt.Println("Servidor rodando na porta:", config.WebServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", config.WebServerPort), r)
}
