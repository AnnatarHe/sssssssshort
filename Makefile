release: export GOOS=linux
release: export GOARCH=amd64

release:
	CGO_ENABLED=1 go build -tags release -o up-clippingkk-api main.go
