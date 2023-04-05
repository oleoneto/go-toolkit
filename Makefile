GOBIN := $(GOPATH)/bin
TARGET_FILE = go-toolkit

clean:
	rm -rf $(TARGET_FILE)

clean-test:
	@go fmt ./...
	@go clean -testcache

build-deps:
	@go mod tidy

test: clean-test
	go test -cover -coverprofile=coverage.out -p 1 ./... | tee test.log
	go tool cover -html=coverage.out

build: build-deps
	@go build -o $(TARGET_FILE)

install: build
	@go env -w GOBIN=$(GOBIN)
	@go install
