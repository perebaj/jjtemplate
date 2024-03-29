# TOOLS VERSIONS
GO_VERSION=1.21.5
GOLANGCI_LINT_VERSION=v1.55.2
devimage={{.Name}}-dev
gopkg=$(devimage)-gopkg
gocache=$(devimage)-gocache
devrun=docker run --rm \
	-v `pwd`:/app \
	-v $(gopkg):/go/pkg \
	-v $(gocache):/root/.cache/go-build \
	$(devimage)
image={{.Registry}}
version=$(shell git rev-parse --short HEAD)

## run all tests. Usage `make test` or `make test testcase="TestFunctionName"` to run an isolated tests
.PHONY: test
test:
	if [ -n "$(testcase)" ]; then \
		go test ./... -timeout 10s -race -run="^$(testcase)$$" -v; \
	else \
		go test ./... -timeout 10s -race; \
	fi

## Calculate code coverage
.PHONY: coverage
coverage:
	go test ./... -coverprofile=c.out 
	go tool cover -html=c.out

## builds the service
.PHONY: service
service:
	go build -o ./cmd/{{.Name}}/{{.Name}} ./cmd/{{.Name}}

## runs the service locally
.PHONY: run
run: service
	./cmd/{{.Name}}/{{.Name}}

## lint the whole project
.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...


## Build the service image
.PHONY: image
image:
	docker build . \
		--build-arg GO_VERSION=$(GO_VERSION) \
		-t $(image)

## Publish the service image
.PHONY: image/publish
image/publish: image
	docker push $(image)

.PHONY: dev
dev: dev/image
	$(devrun)

## Create the dev container image
.PHONY: dev/image
dev/image:
	docker build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION) \
		-t $(devimage) \
		-f Dockerfile.dev \
		.

##run a make target inside the dev container
dev/%: dev/image
	$(devrun) make ${*}

## Display help for all targets
.PHONY: help
help:
	@awk '/^.PHONY: / { \
		msg = match(lastLine, /^## /); \
			if (msg) { \
				cmd = substr($$0, 9, 100); \
				msg = substr(lastLine, 4, 1000); \
				printf "  ${GREEN}%-30s${RESET} %s\n", cmd, msg; \
			} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
