PWD = $(shell pwd)

.PHONY: run

.PHONY: run
run: ## run in local
	go mod tidy
	CGO_ENABLED=1 go run .

.PHONY: build
build: ## build
	goreleaser --skip-validate --snapshot --skip-publish --rm-dist

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
