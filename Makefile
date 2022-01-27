PROJECT_NAME := "go-use"
PKG := "github.com/halpdesk/$(PROJECT_NAME)"
PKG_LIST := $(shell go1.18beta1 list ./... | grep -v /vendor/)

# @golangci-lint run
lint: ## lint files
	@revive -config config.toml -formatter friendly -exclude /vendor/ ./...
test: ## run tests
	@go1.18beta1 test -cover -short -v ${PKG_LIST}
help: ## display help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
