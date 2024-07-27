prot:
	protoc ./proto/*.proto --go_out=. --go-grpc_out=.
build-css:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
run:
	go run main.go