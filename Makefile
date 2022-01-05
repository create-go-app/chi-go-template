.PHONY: clean critic security lint test build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

docker.run: docker.network docker.chi

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.chi.build:
	docker build -t chi .

docker.chi: docker.chi.build
	docker run --rm -d \
		--name cgapp-chi \
		--network dev-network \
		-p 5000:5000 \
		chi

docker.stop: docker.stop.chi

docker.stop.chi:
	docker stop cgapp-chi
