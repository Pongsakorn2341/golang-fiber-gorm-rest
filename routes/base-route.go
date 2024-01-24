package routes

import "github.com/gofiber/fiber/v2"

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func SetUpBaseRoutes(api fiber.Router) {
	// welcome endpoint
	api.Get("/", welcome)
	setUpUserRoute(api)
	setUpProductRoute(api)
	setUpOrderRoutes(api)
}

func setUpUserRoute(api fiber.Router) {
	userApi := api.Group("/users")

	userApi.Post("/", CreateUser)
	userApi.Get("/", GetUsers)
	userApi.Get(":id", GetUser)
	userApi.Put(":id", UpdateUser)
	userApi.Delete("/:id", DeleteUser)
}
func setUpProductRoute(api fiber.Router) {
	productApi := api.Group("/products")

	productApi.Post("/", CreateProduct)
	productApi.Get("/", GetProducts)
	productApi.Get(":id", GetProduct)
	productApi.Put(":id", UpdateProduct)
	productApi.Delete(":id", DeleteProduct)
}

func setUpOrderRoutes(api fiber.Router) {
	orderApi := api.Group("/orders")

	orderApi.Post("", CreateOrder)
	orderApi.Get("/", GetOrders)
	orderApi.Get(":id", GetOrder)
	// orderApi.Put(":id", UpdateProduct)
	// orderApi.Delete(":id", DeleteProduct)
}
