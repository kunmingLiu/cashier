test:
	go test ./...

test-no-cache:
	go clean -testcache && make test

generate:
	go generate ./...