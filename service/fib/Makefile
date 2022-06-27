name     = fibonacci-service
revision ?= $(shell git describe --tags --long --dirty --always)

export AWS_DEFAULT_REGION = $(region)

default: help
.PHONY: default

help: # This
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	  awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
.PHONY: help

binary: ## Build Linux binary ready for containerisation
	$(call cyan, "Building Linux binary ready for containerisation...")
	$(call check-dependency,docker)
	docker run --rm --interactive --tty \
	  --volume "$(shell pwd)":/go/src/$(name) \
	  --workdir /go/src/$(name) \
	  golang:alpine \
	  go build -ldflags  "\
	    -w -s \
	    -X main.revision=${revision} \
	" \
	-o app ./cmd/http
.PHONY: binary

image: binary ## Build docker image
	$(call cyan, "Building docker image...")
	$(call check-dependency,docker)
	docker build --tag $(name):latest .
	@$(MAKE) clean
.PHONY: image

run: image ## Run Docker image locally
	$(call cyan, "Running Docker image locally...")
	$(call check-dependency,docker)
	docker run --rm --interactive --tty --publish 8001:8080 --env-file .env $(name)
.PHONY: run

lint:
	$(call cyan, "Linting...")
	$(call check-dependency,golint)
	@golint ./... | grep -v unexported || true
	@go vet ./... 2>&1 || echo ''
.PHONY: lint

test: lint
	$(call cyan, "Testing...")
	@go test -cover -v ./...
.PHONY: test

start: test ## Run locally
	$(call cyan, "Running...")
	@sh -c "set -a && source .env && set +a && \
		go run -ldflags '\
	    	-w -s \
	    	-X main.revision=${revision} \
	  	' \
	  	./cmd/http \
	"
.PHONY: start

watch: ## Run locally and monitor for changes
	$(call check-dependency,entr)
	@find . \( -path './vendor/*' -or -path './scripts/*' \) -prune -false -o -name '*.go' \
	  | entr -rcs '$(MAKE) start'
.PHONY: dev

clean: 
	@rm -f app
	@go clean
	@git clean --dry-run
.PHONY: clean

amqp:
	@bash -c "set -a && source .env && set +a && \
		go run ./cmd/amqp \
	"
.PHONY: amqp

define check-dependency
	$(if $(shell command -v $(1)),,$(error Make sure $(1) is installed))
endef

define cyan
	@tput setaf 6
	@echo $1
	@tput sgr0
endef
