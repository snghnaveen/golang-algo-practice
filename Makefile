.PHONY: run-all-tests
run-all-tests: clear-cache
		go test -v ./...

.PHONY: clear-cache
clear-cache:
		go clean -testcache

