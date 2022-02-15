.PHONY: build
build:
	go build -o urlShortener/urlShortener -v ./urlShortener/cmd/urlShortener/main.go

image:
	docker build -t urlshortener:1.0 .

.DEFAULT_GOAL := build