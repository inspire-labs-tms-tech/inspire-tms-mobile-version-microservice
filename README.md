# Version Microservice

A microservice to retrieve the public (released) version of an application from the
Apple App Store and Google Play Store.

[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?repo=https://github.com/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice/tree/main)

## API Documentation

https://inspire-labs-tms-tech.github.io/inspire-tms-mobile-version-microservice/


## Caching

Requests are cached with a 10-minute expiration time to avoid spam and flagging by Apple and Google.


## Deploying

> Note: Change the bundles `com.example.*` below to match your application

```shell
docker run \
  --rm \
  --name mobile-app-version-microservice \
  -p 8080:8080 \
  -e APPLE_BUNDLE_ID=com.example.apple \
  -e GOOGLE_APP_ID=com.example.google \
  ghcr.io/inspire-labs-tms-tech/inspire-tms-mobile-version-microservice:latest
```