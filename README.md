# Version Microservice

A microservice to retrieve the public (released) version of an application from the
Apple App Store and Google Play Store.

## Setup

## Build

```shell
docker build \
    --build-arg APPLE_BUNDLE_ID=com.example.apple \
    --build-arg GOOGLE_APP_ID=com.example.google \
    -t my-go-app .
```