.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: deps
deps:
	go mod tidy

.PHONY: test
test:
	go test -cover -count=1 ./...

.PHONY: lint
lint:
	golangci-lint run -c golangci.yaml

.PHONY: build
build:
	env GOOS=linux GOARCH=arm64 go build -o hello
