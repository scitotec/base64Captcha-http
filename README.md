# base64captcha-http

A HTTP API for generating Captchas with [mojocn/base64Captcha](https://github.com/mojocn/base64Captcha).


## Development

### Pushing a new docker image

```bash
docker buildx build \
    --platform linux/amd64,linux/arm64 \
    --push \
    -t scitotec/base64captcha-http:latest \
    -t scitotec/base64captcha-http:1.0.0 \
    -t scitotec/base64captcha-http:1.0.0-base64captcha-1.3 \
    -t scitotec/base64captcha-http:1.0 \
    -t scitotec/base64captcha-http:1.0-base64captcha-1.3 \
    -t scitotec/base64captcha-http:1-base64captcha-1.3 \
    -t scitotec/base64captcha-http:1 \
    .
```