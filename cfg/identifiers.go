package cfg

import "os"

var AppleBundleId = os.Getenv("APPLE_BUNDLE_ID")
var GoogleAppId = os.Getenv("GOOGLE_APP_ID")
