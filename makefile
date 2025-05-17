APP_NAME=blog-go
MAIN=./cmd/main.go

dev:
	air

templ:
	templ generate

build:
	CGO_ENABLED=0 GOOS=linux go build -o $(APP_NAME) $(MAIN)

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run --rm -p 8080:8080 $(APP_NAME):latest

clean:
	rm -f $(APP_NAME)

build-all: templ build

docker-prod:
	docker build -f Dockerfile -t $(APP_NAME):prod .
