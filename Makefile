.PHONY: build
build:
	go build -o urlShortener/urlShortener -v ./urlShortener/cmd/urlShortener/main.go

image: build
	docker build -t test .
run: image
	docker run -dit --name sql -p 5432:5432 test

clean:
	docker stop $$(docker ps -a -q)
	docker system prune -a
exec:
	docker exec -it sql bash


.DEFAULT_GOAL := build
