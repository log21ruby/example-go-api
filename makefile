dev: 
	nodemon --exec go run cmd/api/main.go --signal SIGTERM

doc: 
	swag init -g ./cmd/api/main.go
	
run-test:
	go test ./...

run:
	go run main.go

build:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker build . -t ${tag}