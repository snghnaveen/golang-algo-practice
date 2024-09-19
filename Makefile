.PHONY: tests
tests: clear-cache
	go test -v ./...

.PHONY: clear-cache
clear-cache:
	go clean -testcache

.PHONY: gofmt
gofmt:
	go fmt ./...

.PHONY: docker-build
docker-build:
	docker build . -t golang-algo-practice

.PHONY: docker-run
docker-run: docker-build
	docker rm -f golang-algo-practice || true
	docker run --name=golang-algo-practice golang-algo-practice