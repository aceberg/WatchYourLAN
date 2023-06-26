DUSER=aceberg
DNAME=watchyourlan

IFACE=enp4s0
DBPATH=/data/hosts.db
SHOUTRRR_URL=gotify://192.168.2.1:8083/AwQqpAae.rrl5Ob/?title=Unknown host detected&DisableTLS=yes

mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/WatchYourLAN && \
	go mod tidy

run:
	cd cmd/WatchYourLAN/ && \
	sudo \
	env IFACE=$(IFACE) DBPATH=$(DBPATH) THEME=$(THEME) \
	go run . #-n http://192.168.2.3:8850

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/WatchYourLAN/ && \
	CGO_ENABLED=0 go build -o ../../tmp/WatchYourLAN .

docker-build:
	docker build -t $(DUSER)/$(DNAME) .

docker-run:
	docker rm wyl || true
	docker run --name wyl \
	-e "IFACE=$(IFACE)" \
	-e "SHOUTRRR_URL=$(SHOUTRRR_URL)" \
	-e "TZ=Asia/Novosibirsk" \
	--network="host" \
	-v ~/.dockerdata/wyl:/data \
	$(DUSER)/$(DNAME)

clean:
	rm src/$(DNAME) || true
	docker rmi -f $(DUSER)/$(DNAME)

dev: docker-build docker-run