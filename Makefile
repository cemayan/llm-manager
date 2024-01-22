default: help

PROJECT_NAME=llm-manager
ARCH=$(arch)
CLI_MAIN_FOLDER=./cmd
BIN_FOLDER=bin
BIN_FOLDER_MACOS=${BIN_FOLDER}/${ARCH}/darwin
BIN_FOLDER_WINDOWS=${BIN_FOLDER}/${ARCH}/windows
BIN_FOLDER_LINUX=${BIN_FOLDER}/${ARCH}/linux
BIN_FOLDER_SCRATCH=${BIN_FOLDER}/${ARCH}/scratch

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent
# LDFLAGS=-X main.buildDate=`date -u +%Y-%m-%dT%H:%M:%SZ` -X main.version=`scripts/version.sh`
LDFLAGS=

## setup: install all build dependencies for ci
setup: mod-download

## compile-macos: compiles project for macos
compile-macos: fmt test build-macos

## compile-windows: compiles project for windows
compile-windows: fmt test build-windows

## compile-linux: compiles project for linux
compile-linux: fmt test build-linux

## compile-alpine: compiles project for alpine-scratch
compile-alpine-scratch: fmt test build-alpine-scratch

skaffold-dev:
	@echo "  >  Building binary for MacOS"
	skaffold dev

build:
	@echo "  >  Building binary"
	 && go build \
		-ldflags="${LDFLAGS}" \
		-o ${BIN_FOLDER}/${BIN_NAME} \
		"${CLI_MAIN_FOLDER}"

build-all: build-macos build-windows build-linux build-alpine-scratch

build-macos:
	@echo "  >  Building binary for MacOS"
	 GOOS=darwin GOARCH=${ARCH} \
		go build \
		-ldflags="${LDFLAGS}" \
		-o ${BIN_FOLDER_MACOS}/${BIN_NAME} \
		"${CLI_MAIN_FOLDER}"

build-windows:
	@echo "  >  Building binary for Windows"
	 GOOS=windows GOARCH=${ARCH} \
		go build \
		-ldflags="${LDFLAGS}" \
		-o ${BIN_FOLDER_WINDOWS}/${BIN_NAME}.exe \
		"${CLI_MAIN_FOLDER}"

build-linux:
	@echo "  >  Building binary for Linux"
	 GOOS=linux GOARCH=${ARCH} \
		go build \
		-ldflags="${LDFLAGS}" \
		-o ${BIN_FOLDER_LINUX}/${BIN_NAME} \
		"${CLI_MAIN_FOLDER}"

# Alpine & scratch base images use musl instead of gnu libc, thus we need to add additional parameters on the build
build-alpine-scratch:
	@echo "  >  Building binary for Alpine/Scratch"
	 CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} \
		go build \
		-ldflags="${LDFLAGS}" \
		-a -installsuffix cgo \
		-o ${BIN_FOLDER_SCRATCH}/${BIN_NAME} \
		"${CLI_MAIN_FOLDER}"

fmt:
	@echo "  >  Formatting code"
	go fmt ./...

mod-download:
	@echo "  >  Download dependencies..."
	go mod download && go mod tidy

test:
	@echo "  >  Executing unit tests"
	go test -v -cover  ./...


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo