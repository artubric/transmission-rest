# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-alpine AS build

ARG BUILD_OS_TARGET
ARG BUILD_ARCH_TARGET

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY config ./config
COPY internal ./internal

RUN GOOS=$BUILD_OS_TARGET GOARCH=$BUILD_ARCH_TARGET CGO_ENABLED=0 go build -o /transmission-rest ./cmd/transmission-rest/main.go

CMD [ "/transmission-rest" ]

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /transmission-rest /transmission-rest
COPY --from=build /app/config/.env.local /app/config/.env.prod /config/

ENTRYPOINT ["/transmission-rest"]