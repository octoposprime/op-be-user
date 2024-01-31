GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
BINARY_NAME=main
VERSION?=0.0.0#Docker image release version
DOCKER_REGISTRY?=ghcr.io#Docker registry
DOCKER_REPOSITORY?=octoposprime#Docker repository owner
DOCKER_CONTAINER?=op-be-user#Docker image name
EXPORT_RESULT?=false# for CI please set EXPORT_RESULT to true
TEST?=false#is the container test?
POSTGRES_USERNAME?=op#Postgres Db User Name
POSTGRES_PASSWORD?=op#Postgres Db Password
JWT_SECRET_KEY?=op#Jwt Secret Key
REDIS_PASSWORD?=op#Redis Password
LOCAL_PORT=18082#Grpc port for local
CONTAINER_PORT=18080#Grpc port in container
NETWORK=op#Docker network name

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build coverage lint

all: help

## Build:
build: ## Build your project and put the output binary in out/bin/
	mkdir -p out/bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o out/bin/$(BINARY_NAME) ./cmd/service/main.go

clean: ## Remove build related file
	rm -rf out/

## Test:
test: ## Run the tests of the project
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

coverage: ## Run the tests of the project and export the coverage
	$(GOTEST) -cover -covermode=count -coverprofile=out/profile.cov ./...
	$(GOCMD) tool cover -func out/profile.cov
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/AlekSi/gocov-xml
	GO111MODULE=off go get -u github.com/axw/gocov/gocov
	gocov convert out/profile.cov | gocov-xml > out/coverage.xml
endif

## Lint:
lint: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > out/checkstyle-report.xml" || echo "" ))
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --deadline=65s $(OUTPUT_OPTIONS)

## Docker:
docker-build: ## Use the dockerfile to build the container
	docker build --rm --tag $(DOCKER_REPOSITORY)/$(DOCKER_CONTAINER) \
	--build-arg TEST=$(TEST) \
	--build-arg POSTGRES_USERNAME=$(POSTGRES_USERNAME) \
	--build-arg POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	--build-arg JWT_SECRET_KEY=$(JWT_SECRET_KEY) \
	--build-arg REDIS_PASSWORD=$(REDIS_PASSWORD) .

docker-release: ## Release the container with tag latest and version
	docker tag $(DOCKER_REPOSITORY)/$(DOCKER_CONTAINER) $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY)/$(DOCKER_CONTAINER):latest
	docker tag $(DOCKER_REPOSITORY)/$(DOCKER_CONTAINER) $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY)/$(DOCKER_CONTAINER):$(VERSION)
	# Push the docker images
	docker push $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY)/$(DOCKER_CONTAINER):latest
	docker push $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY)/$(DOCKER_CONTAINER):$(VERSION)

## Run:
local-run: ## Run in Local for Development
	LOCAL=true go run cmd/service/main.go

docker-run: ## Run in Docker for Development
	docker run -d --expose $(LOCAL_PORT) -p $(LOCAL_PORT):$(CONTAINER_PORT) --network $(NETWORK) --name $(DOCKER_CONTAINER) $(DOCKER_REPOSITORY)/$(DOCKER_CONTAINER)

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)