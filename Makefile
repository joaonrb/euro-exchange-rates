PROJECT_NAME := "$(shell head -1 go.mod | awk '{print $$2}')"
PKG := "$(PROJECT_NAME)"
#PKG := "gitlab.com/joaonrb/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v mocks)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v mocks | grep -v _test.go)

.PHONY: all dep build clean test coverage coverhtml lint

all:
	build

lint: ## Lint the files
	@golint -set_exit_status "${PKG}/..."

fmt: ## Fmt the code
	@go fmt ./...

test: ## Run unittests
	@go test `go list ./... | grep -v mocks`
#	@go test -short ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: dep ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	@rm -fr cover;
	@mkdir cover;
	@for package in ${PKG_LIST}; do \
		go test -covermode=count -coverprofile "cover/$${package##*/}.cov" "$$package"; \
	done;
	@echo "mode: count" > cover/coverage.cov;
	@tail -q -n +2 cover/*.cov >> cover/coverage.cov;
	@go tool cover -func=cover/coverage.cov;


coverhtml: coverage ## Generate global code coverage report in HTML
	@go tool cover -html=cover/coverage.cov -o coverage.html;

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file
	@go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
