.BINARY: wnc-node
GO = go

build:
	$(GO) build -o build/wnc-node ./cmd/wnc-node

clean:
	rm -rf build/

run: build
	./build/wnc-node

release-linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o build/wnc-node-linux-amd64 ./cmd/wnc-node

release-windows:
	GOOS=windows GOARCH=amd64 $(GO) build -o build/wnc-node-windows-amd64.exe ./cmd/wnc-node

release-macos:
	GOOS=darwin GOARCH=arm64 $(GO) build -o build/wnc-node-macos-arm64 ./cmd/wnc-node
