.PHONY: build clean tool lint help

all: build

build:
	@go build -v .

build-upx:
	@rm -rf main
	@go build -ldflags="-s -w" -o main main.go && upx -9 main; true

vet:
	go vet ./pkg/...; true
	go vet ./test/...; true

lint:
	golint ./pkg/...

clean:
	rm -rf service
	rm -rf manager
	go clean -i .

# go install mvdan.cc/gofumpt@latest
fmt:
	@gofumpt -l -w ./pkg/
	@gofumpt -l -w ./test/
	# go fmt ./pkg/...
	# find ./pkg -name "*.go" | xargs gofumpt -l -w

# go install github.com/xxjwxc/gormt@master
gen-model:
	@gormt

gen-proto:
	@cd pkg/server/grpc/gateway && protoc --go_out=plugins=grpc:. server.proto

git-push:
	./cmd.sh git push

git-clear:
	./cmd.sh git clear

install:
	@docker plugin install  grafana/loki-docker-driver:latest --alias loki --grant-all-permissions

upgrade-go:
	@wget https://go.dev/dl/go1.18.4.linux-amd64.tar.gz
	@sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz
	@rm -rf go1.18.4.linux-amd64.tar.gz

help:
	@echo "make: compile packages and dependencies"
	@echo "make vet: run specified go vet"
	@echo "make lint: golint ./..."
	@echo "make fmt: gofumpt -l -w ./pkg/"
	@echo "make clean: remove object files and cached files"

