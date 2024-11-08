default: help
.PHONY: default

name    ?= $(notdir $(realpath $(dir $(realpath $(MAKEFILE_LIST)))))
registry = 123456789012.dkr.ecr.eu-west-2.amazonaws.com
revision ?= $(shell ./scripts/version.sh)

help: # This
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
.PHONY: help

deps: ## Install dependencies
	@echo "==> Installing dependencies"
	@dotnet restore ./FibonacciHttp/FibonacciHttp.csproj
.PHONY: deps

gen: clean ## Run the code generation
	@echo $(revision) > .version
.PHONY: gen

lint: gen ## Run the linter
	$(call cyan, "Linting...")
	@@dotnet build ./FibonacciHttp/FibonacciHttp.csproj /p:TreatWarningsAsErrors=true 2>&1 || echo ''
.PHONY: lint

quality: ## Run various code quality assurance checks
	$(call cyan, "Scanning...")
	$(call check-dependency,docker)
	@docker run --rm -i hadolint/hadolint < Dockerfile 2>&1  || echo ''
.PHONT: quality

image: lint ## Create Docker image
	$(call cyan, "Continerising...")
	$(call check-dependency,docker)
	@docker build --tag $(name):latest --tag $(name):$(revision) .
.PHONY: image

publish: image ## Publish to container registory, assumes image is built
	$(call cyan, "Publishing...")
	$(call check-dependency,docker)
	$(call check-dependency,aws)
	@docker tag $(name):latest $(registry)/$(name):$(revision)
	@aws ecr get-login-password --region eu-west-2 | \
	  docker login --username AWS --password-stdin ${registry}
	@docker push $(registry)/$(name):$(revision)
.PHONY: publish

test: lint ## Run the project tests
	$(call cyan, "Testing...")
	$(call setenv,)
	@dotnet test ./FibonacciHttp/FibonacciHttp.csproj --no-build --logger "console;verbosity=detailed"
.PHONY: test

start: test ## Start the server
	$(call cyan, "Running...")
	$(call setenv,)
	@dotnet run --project ./FibonacciHttp/FibonacciHttp.csproj
.PHONY: start

clean: 
	@dotnet clean ./FibonacciHttp/FibonacciHttp.csproj
	@rm -fr $(output)
.PHONY: clean

define check-dependency
	$(if $(shell command -v $(1)),,$(error Make sure $(1) is installed))
endef
