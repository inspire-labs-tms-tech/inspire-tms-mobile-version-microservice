package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice/cfg"
	"github.com/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice/internal/routes/responses"
	"github.com/n0madic/google-play-scraper/pkg/app"
)

func GetGooglePlayStoreVersion(c *fiber.Ctx) error {

	a := app.New(cfg.GoogleAppId, app.Options{
		Country:  "us",
		Language: "us",
	})

	err := a.LoadDetails()
	if err != nil {
		return c.JSON(responses.NewErrorResponse(err.Error()))
	}

	return c.JSON(responses.NewVersionResponse(a.Version))

}
