.PHONY: client-build client-docker-all client-docker-build client-docker-push

# go commands
GOCMD=go
GOBUILD=$(GOCMD) build

# dist
BIN_DIR=bin
REGISTRY_BASE=ghcr.io/inonius
DOCKER_IMG_TAG ?= $(or $(INONIUS_DOCKER_IMG_TAG),develop)

V3CLI_IMG := $(REGISTRY_BASE)/client:$(DOCKER_IMG_TAG)

client-build: ## single binary build
	$(GOBUILD)  -o ./$(BIN_DIR)/inonius_v3cli main.go

client-docker-all: client-docker-build client-docker-push ## docker build and push

client-docker-build:
	docker build \
		--file Dockerfile \
		--tag ${V3CLI_IMG} .

client-docker-push: ## docker push
	docker push  ${V3CLI_IMG}
