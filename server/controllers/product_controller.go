package controllers

import (
	"context"
	"log"

	"github.com/Nahuel-199/go-crud.git/database"
	Models "github.com/Nahuel-199/go-crud.git/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(ctx *fiber.Ctx) error {
	// Crea un nuevo producto utilizando los datos de la solicitud HTTP
	product := new(Models.Product)
	if err := ctx.BodyParser(product); err != nil {
		return err
	}

	// Inserta el nuevo producto en la base de datos
	collection := database.GetCollection("products")
	result, err := collection.InsertOne(context.Background(), product)
	if err != nil {
		return err
	}

	// Devuelve una respuesta con el ID del nuevo producto insertado
	return ctx.JSON(fiber.Map{
		"message": "Producto creado",
		"id":      result.InsertedID,
	})
}

func GetAllProducts(c *fiber.Ctx) error {
	//Establecemos la conexion con la base de datos
	collection := database.GetCollection("products")

	// Realizamos la consulta para traer todos los productos
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// Cerramos el cursor al terminar
	defer cursor.Close(context.Background())

	// Creamos un slice para almacenar los resultados de la consulta
	var products []Models.Product

	// Iteramos sobre los resultados y los almacenamos en el slice
	for cursor.Next(context.Background()) {
		var product Models.Product
		err := cursor.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	// Devolvemos los productos en formato JSON
	return c.JSON(products)

}

func GetProductById(c *fiber.Ctx) error {
	//Obtener el ID del producto de la URL
	id := c.Params("id")

	collection := database.GetCollection("products")

	//Convertir el ID en un objectID de MongoDB
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	//Obtener el producto de la base de datos
	product := &Models.Product{}
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(product)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Producto no encontrado",
		})
	}

	// Devolver el producto como respuesta
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	//Obtener el ID del producto de la URL
	id := c.Params("id")

	//Convertirmos el ID en un ObjectId de mongoDB
	objID, err := primitive.ObjectIDFromHex(id)

	//Si hay un error o si es igual a null
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	//Creamos un objeto de tipo Product y lo enlazamos con los datos enviados en la solicitud JSON
	var product Models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	//Apuntamos a la coleccion donde vamos a actualizar
	collection := database.GetCollection("products")

	filter := bson.M{"_id": objID}

	// Creamos un objeto de tipo bson.M con los nuevos valores para actualizar el producto
	update := bson.M{
		"$set": bson.M{
			"title":       product.Title,
			"description": product.Description,
			"price":       product.Price,
			"img":         product.Img,
			"inStock":     product.InStock,
		},
	}

	// Actualizamos el producto en la base de datos
	if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
		return err
	}

	// Devolvemos una respuesta con estado 200 (OK) y un mensaje de éxito
	return c.SendString("Product updated successfully!")
}

func DeleteProduct(c *fiber.Ctx) error {
	// Obtenemos el id del producto a eliminar de los parámetros de la ruta
	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := database.GetCollection("products")

	filter := bson.M{"_id": objID}

	// Eliminamos el producto de la base de datos
	if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	// Devolvemos una respuesta con estado 200 (OK) y un mensaje de éxito
	return c.SendString("Product deleted successfully!")
}
