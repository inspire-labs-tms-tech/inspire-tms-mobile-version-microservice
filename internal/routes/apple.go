package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice/cfg"
	"github.com/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice/internal/routes/responses"
	"io"
	"net/http"
)

type iTunesResponse struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		Version string `json:"version"`
	} `json:"results"`
}

func GetAppleAppStoreVersion(c *fiber.Ctx) error {

	var url = fmt.Sprintf("https://itunes.apple.com/lookup?bundleId=%s", cfg.AppleBundleId)
	resp, err := http.Get(url)

	if err != nil {
		return c.JSON(responses.NewErrorResponse(err.Error()))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(responses.NewErrorResponse(err.Error()))
	}

	var data iTunesResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return c.JSON(responses.NewErrorResponse(err.Error()))
	}

	if data.ResultCount == 0 {
		return c.JSON(responses.NewErrorResponse("no results found"))
	}

	return c.JSON(responses.NewVersionResponse(data.Results[0].Version))

}
