# Makefile for the OrbundAwardIntegrationJob batch job

main_package_path =./
binary_name = fsa_cli
build_dir = ./bin

.PHONY: help
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo 'Are you sure [y/N] ' && read ans && [ $${ans:-N} = y ]

## build: Build the application for MacOS, Linux, and Windows
.PHONY: build
build: lint
	@mkdir -p ${build_dir}
	GOARCH=amd64 GOOS=darwin go build -o ${build_dir}/${binary_name}-darwin-amd64 ${main_package_path}
	GOARCH=arm64 GOOS=darwin go build -o ${build_dir}/${binary_name}-darwin-arm64 ${main_package_path}
	GOARCH=amd64 GOOS=linux go build -o ${build_dir}/${binary_name}-linux ${main_package_path}
	GOARCH=amd64 GOOS=windows go build -o ${build_dir}/${binary_name}-windows-amd64.exe ${main_package_path}

## clean: Clean up the build binaries
.PHONY: clean
clean: confirm
	@echo "Cleaning up..."
	@rm -rf ${build_dir}

## fmt: Format the code using gofmt
fmt:
	go fmt ./...

## vet: Vet the code using go vet
vet: fmt
	go vet ./...

## lint: Lint the code using golangci-lint
lint: vet
	@golangci-lint run
