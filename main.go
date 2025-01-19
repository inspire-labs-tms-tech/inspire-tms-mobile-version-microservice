package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice/internal/middleware"
	"github.com/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice/internal/routes"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

func main() {
	app := fiber.New()

	localCache := cache.New(10*time.Minute, 20*time.Minute) // setting default expiration time and clearance time.

	app.Get("/ping", routes.GetPing)
	app.Get("/apple", middleware.CacheMiddleware(localCache), routes.GetAppleAppStoreVersion)
	app.Get("/google", middleware.CacheMiddleware(localCache), routes.GetGooglePlayStoreVersion)

	err := app.Listen(":8080")
	if err != nil {
		log.Printf(err.Error())
	}
}
