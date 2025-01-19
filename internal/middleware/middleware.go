package middleware

import (
	"encoding/json"
	"github.com/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice/internal/routes/responses"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func CacheMiddleware(cache *cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Only cache GET requests
		if c.Method() != "GET" {
			return c.Next()
		}

		cacheKey := c.Path()

		// Check if the response is already in the cache
		if cached, found := cache.Get(cacheKey); found {
			c.Response().Header.Set("Cache-Status", "HIT")
			return c.JSON(cached)
		}

		c.Set("Cache-Status", "MISS")
		err := c.Next()
		if err != nil {
			return err
		}

		var data responses.VersionResponse

		body := c.Response().Body()
		err = json.Unmarshal(body, &data)
		if err != nil {
			return c.JSON(responses.NewErrorResponse(err.Error()))
		}

		// Cache the response for 10 minutes
		cache.Set(cacheKey, data, 10*time.Minute)

		return nil
	}
}
