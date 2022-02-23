build:
	go build -o urlShortener -v ./cmd/urlShortener/main.go
	@echo "\033[0;32mBinary built successfully\033[m"
test_inmem:
	 go test -v -count=1 internal/app/apiserver/apiserver_inmem_GetPost_test.go internal/app/apiserver/apiserver.go
test_psql:
	 go test -v -count=1 internal/app/apiserver/apiserver_psql_GetPost_test.go internal/app/apiserver/apiserver.go
run_inmem:
	./urlShortener
run_psql:
	./urlShortener -mem psql
