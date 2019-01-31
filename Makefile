.PHONY: build

build:
	@go mod vendor
	@go vet ./...
	@go install ./...

