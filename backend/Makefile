mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/WatchYourLAN && \
	go mod tidy

run:
	cd cmd/WatchYourLAN/ && \
	sudo \
	go run . #-n http://192.168.2.3:8850

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

swag:
	swag init -g main.go --dir cmd/WatchYourLAN,internal/api,internal/models