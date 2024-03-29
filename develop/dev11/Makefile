all: help

.PHONY: deps
# Update dependencies
deps: init
	go mod tidy -v

.PHONY: lint
# Lint the project
lint:
	golangci-lint run ./src/...

.PHONY: run
# Run the application and the database container
run:
	go run ./src/cmd/web

.PHONY: run-air
# Run the application in live reload mode
run-air:
	air

.PHONY: init
# Initialize the repository for development
init: install-gofumpt install-air install-golangci-lint

.PHONY: install-gofumpt
install-gofumpt:
ifeq (, $(shell which gofumpt))
	echo "Installing gofumpt..."
	go install mvdan.cc/gofumpt@latest
endif

.PHONY: install-air
install-air:
ifeq (, $(shell which air))
	echo "Installing air..."
	go install github.com/cosmtrek/air@latest
endif

.PHONY: install-golangci-lint
install-golangci-lint:
ifeq (, $(shell which golangci-lint))
	echo "Installing golangci-lint..."
	$(shell curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.55.2)
endif

.PHONY: help
# Show this help
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t
