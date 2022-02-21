build:
	go build -o urlShortener -v ./cmd/urlShortener/main.go
	@echo "\033[0;32mBinary built successfully\033[m"
test:
	go test -v ./internal/app/apiserver/apiserver_Post_test.go
run_inmem:
	./urlShortener
run_psql:
	./urlShortener -mem psql
