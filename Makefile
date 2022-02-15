.PHONY: build
build:
	go build -v ./urlShortener/cmd/urlShortener

.DEFAULT_GOAL := build