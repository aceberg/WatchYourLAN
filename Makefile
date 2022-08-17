DUSER=aceberg
DNAME=watchyourlan
DPORT=8840
VERSION=0.4

mod:
	rm go.mod
	go mod init $(DNAME)
	go mod tidy

run:
	sudo go run .

go-build:
	go build .

docker-build:
	docker build -t $(DUSER)/$(DNAME):latest -t $(DUSER)/$(DNAME):$(VERSION) .

docker-run:
	docker rm wyl || true
	docker run --name wyl \
	-e "IFACE=virbr-bw" \
	--network="host" \
	-v ~/.dockerdata/wyl:/data \
	$(DUSER)/$(DNAME):$(VERSION)

clean:
	rm $(DNAME) || true
	docker rmi -f $(DUSER)/$(DNAME):latest $(DUSER)/$(DNAME):$(VERSION)

dev: go-build docker-build docker-run