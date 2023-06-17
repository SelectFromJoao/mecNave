package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-TOKEN",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	SetupUsersRoutes(v1)
	SetupCompanyRoutes(v1)
	SetupBannersRoutes(v1)
}

func SetupUsersRoutes(v1 fiber.Router) {
	v1.Get("/users/:id", GetUserByID)
	v1.Post("/users", Create)
	v1.Put("/users", Update)
	v1.Delete("/users", Delete)
}

func SetupCompanyRoutes(v1 fiber.Router) {
	v1.Get("/companies/:id", GetCompanyByID)
	v1.Post("/companies", CreateCompany)
	v1.Put("/companies", UpdateComapany)
	v1.Delete("/companies", DeleteCompany)
}

func SetupBannersRoutes(v1 fiber.Router) {
	v1.Get("/banners/:id", GetBannerByID)
	v1.Get("/banners", GetAllBanners)
	v1.Post("/banners", CreateBanner)
	v1.Put("/banners", UpdateComapany)
	v1.Delete("/banners", DeleteBanner)
}
