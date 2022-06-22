FROM golang:1.18-alpine as build

WORKDIR /build

COPY . /build/

RUN go build


FROM golang:1.18-alpine

WORKDIR /app
COPY healthcheck.sh /app/
COPY --from=build /build/base64captcha-http /app/

HEALTHCHECK CMD [ "/app/healthcheck.sh" ]

EXPOSE 8777

ENTRYPOINT ["./base64captcha-http"]