package router

import (
	middlewareHandler "shopping-service-be/modules/middleware/middlewareHandler"
	productHandler "shopping-service-be/modules/product/productHandler"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(api fiber.Router, productHttpHandler productHandler.ProductHttpHandlerService, middlewareHandler middlewareHandler.MiddlewareHandlerService) {
	routes := api.Group("/product")

	routes.Get("/", middlewareHandler.JwtAuthorization, productHttpHandler.GetAllProducts)
	routes.Post("/", productHttpHandler.CreateProduct)
	routes.Put("/:product_id", productHttpHandler.UpdateProductById)
	routes.Delete("/:product_id", productHttpHandler.DeleteProductById)

	routes.Post("/form-data", productHttpHandler.CreateProductFormData)
	routes.Post("/json", productHttpHandler.CreateProductJson)
	routes.Post("/urlencoded", productHttpHandler.CreateProductUrlencoded)

	routes.Post("/upload", productHttpHandler.UploadProductImage)
}
