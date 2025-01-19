package cfg

import (
	"log"
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
		log.Fatal("APPLE_BUNDLE_ID runtime variable not set")
	}

	if GoogleAppId == "" {
		log.Fatal("GOOGLE_APP_ID runtime variable not set")
	}
}
