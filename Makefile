# Makefile
.PHONY: build

BINARY_NAME=go-web-accelerator

# builds the tailwind css sheet, and compiles the binary into a usable thing.
build:
	go mod tidy && \
   	templ generate && \
	go generate && \
	CGO_ENABLED=0 go build -ldflags="-w -s" -o ${BINARY_NAME}

# dev runs the development server where it builds the tailwind css sheet,
# and compiles the project whenever a file is changed.
dev:
	docker compose up -d &\
	templ generate --watch --cmd="go generate" &\
	templ generate --watch --cmd="go run ."

stop:
	docker compose stop &\
	go clean

build:
	docker-compose -p go-web-accelerator build

start:
	docker-compose up -d