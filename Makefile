mod:
	rm go.mod
	go mod init watchyourlan
	go mod tidy
run:
	go run .
build:
	go build .