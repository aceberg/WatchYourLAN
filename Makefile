DUSER=aceberg
DNAME=watchyourlan

VERSION=0.5

mod:
	cd src && \
	rm go.mod && \
	go mod init $(DNAME) && \
	go mod tidy

run:
	cd src && \
	sudo go run .

go-build:
	cd src && \
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
	rm src/$(DNAME) || true
	docker rmi -f $(DUSER)/$(DNAME):latest $(DUSER)/$(DNAME):$(VERSION)

dev: go-build docker-build docker-run