test:
	go test -v ./...

lint:
	golangci-lint run --timeout=5m

ci: lint test
