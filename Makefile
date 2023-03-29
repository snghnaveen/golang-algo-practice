.PHONY: tests
tests: clear-cache
		go test -v ./...

.PHONY: clear-cache
clear-cache:
		go clean -testcache

.PHONY: gofmt
gofmt:
		go fmt ./...

