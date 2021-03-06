APP_NAME = transmission-rest
VERSION = 0.0.1
COMPILE_NAME = ${APP_NAME}-${VERSION}

BUILD_OS_TARGET = $(shell go env GOOS)
BUILD_ARCH_TARGET = $(shell go env GOARCH)

# build local executable
build-local:
	go build -o bin/${COMPILE_NAME} cmd/transmission-rest/main.go

# build docker image (build for local by default)
build-docker-image: set-docker-tag save-docker-image
	docker build -f ./build/Dockerfile . \
		-t ${DOCKER_TAG} \
		--build-arg BUILD_OS_TARGET=${BUILD_OS_TARGET} \
		--build-arg BUILD_ARCH_TARGET=${BUILD_ARCH_TARGET} \

# build docker image for Synology DS920+
build-docker-image-synology-920p: set-synology-920p-vars set-docker-tag build-docker-image save-docker-image

set-synology-920p-vars:
	$(eval BUILD_OS_TARGET=linux)
	$(eval BUILD_ARCH_TARGET=amd64)

set-docker-tag:
	$(eval DOCKER_TAG=${APP_NAME}:${BUILD_OS_TARGET}-${BUILD_ARCH_TARGET}-${VERSION})

save-docker-image:
	docker save -o ./bin/di-${APP_NAME}-${BUILD_OS_TARGET}-${BUILD_ARCH_TARGET}-${VERSION} ${DOCKER_TAG}
