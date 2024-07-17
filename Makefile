include ./deployments/.env

run:
	SERVER_ADDRESS=${SERVER_ADDRESS} \


	go run ./cmd/schedule/main.go
