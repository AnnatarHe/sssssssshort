release: export GOOS=linux
release: export GOARCH=amd64

release:
	CGO_ENABLED=0 go build -tags release -o up-clippingkk-api main.go

.PHONY: db
db:
	docker run --rm -it -p 5432:5432 -e POSTGRES_PASSWORD='admin' -e POSTGRES_DB='clippingkk' postgres

redis:
	docker run --rm -p 6379:6379 redis

deps:
	go get -u github.com/kataras/iris
	go get -u github.com/satori/go.uuid
	go get github.com/jmoiron/sqlx
	go get -u github.com/iris-contrib/middleware/jwt
	go get -u gopkg.in/go-playground/validator.v9
	go get -u github.com/aliyun/alibaba-cloud-sdk-go/services/dm
	go get -u github.com/bsm/redis-lock