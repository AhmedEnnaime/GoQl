IMAGE_NAME := goql

CONTAINER_NAME := goql_ctn

run:
	@docker-compose up -d

test: 
	@go test -v ./...

stop:
	@docker-compose down

