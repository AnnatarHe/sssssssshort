release: export GOOS=linux
release: export GOARCH=amd64

release:
	CGO_ENABLED=1 go build -tags release -o up-clippingkk-api main.go

db:
	docker run -p 3306:3306 --rm -e MYSQL_ROOT_PASSWORD=password -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=short -e MYSQL_USER=short -e MYSQL_PASSWORD=password mysql:5.7