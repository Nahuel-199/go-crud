package main

import (
	"log"

	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/Nahuel-199/go-crud.git/database"

	"github.com/Nahuel-199/go-crud.git/controllers"
)

func main() {

	client, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Crea una nueva instancia de la aplicación Fiber
	app := fiber.New()

	//Rutas
	app.Post("/products", controllers.CreateProduct)

	app.Get("/products", controllers.GetAllProducts)

	app.Get("/product/:id", controllers.GetProductById)

	app.Put("/product/update/:id", controllers.UpdateProduct)

	app.Delete("/product/delete/:id", controllers.DeleteProduct)

	// Inicia la aplicación en el puerto 3000
	log.Fatal(app.Listen(":3000"))
}
