.PHONY: build get

build:
	@go mod vendor
	@go fmt ./...
	@go vet ./...
	@cd stencil && go install -a

