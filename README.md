# Version Microservice

A microservice to retrieve the public (released) version of an application from the
Apple App Store and Google Play Store.

## Using

```shell
docker run \
  --rm \
  --name mobile-app-version-microservice \
  -p 8080:8080 \
  -e APPLE_BUNDLE_ID=com.example.apple \
  -e GOOGLE_APP_ID=com.example.google \
  ghcr.io/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice:latest
```