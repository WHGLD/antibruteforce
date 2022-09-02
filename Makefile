generate:
	protoc --go_out=internal/server --go-grpc_out=internal/server api/abf.proto

build:
	go build -o ./bin/ab_force ./cmd/abf
	go build -o ./bin/migrate ./cmd/migrate
	go build -o ./bin/cli ./cmd/cli

lint:
	golangci-lint run ./...

run:
	docker-compose -f deployments/docker-compose.yaml up

debug:
	docker-compose -f deployments/docker-compose.yaml up -d redis db migrate

down:
	docker-compose -f deployments/docker-compose.yaml down