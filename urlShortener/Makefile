build:
	go build -o urlShortener -v ./cmd/urlShortener/main.go
	@echo "\033[0;32mBinary built successfully\033[m"
test:
	 go test -v -count=1 ./.../apiserver
run_inmem:
	./urlShortener
run_psql:
	./urlShortener -mem psql
