package cfg

import (
	"github.com/fatih/color"
	"os"
)

var (
	AppleBundleId string
	GoogleAppId   string
)

func init() {
	AppleBundleId = os.Getenv("APPLE_BUNDLE_ID")
	GoogleAppId = os.Getenv("GOOGLE_APP_ID")

	if AppleBundleId == "" {
		color.Red("APPLE_BUNDLE_ID runtime variable not set")
		os.Exit(1)
	}

	if GoogleAppId == "" {
		color.Red("GOOGLE_APP_ID runtime variable not set")
		os.Exit(1)
	}
}
