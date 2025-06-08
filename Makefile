BINARY_NAME=permiflow
DIST_DIR=dist

.PHONY: build clean release test

build:
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd
	@echo "✅ Binaries built in $(DIST_DIR)/"

clean:
	rm -rf $(DIST_DIR)
	@echo "🧹 Cleaned up build artifacts"

release:
	goreleaser release --clean

test:
	go test ./...
