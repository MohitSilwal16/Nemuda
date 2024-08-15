prot:
	protoc ./proto/*.proto --go_out=. --go-grpc_out=.
run:
	go run main.go
build:
	env GOOS=linux go build -o server
