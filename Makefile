.PHONY: build clean tool lint help

all: build

build:
	@go build -v .

build-upx:
	@rm -rf main
	@go build -ldflags="-s -w" -o main main.go && upx -9 main; true

vet:
	go vet ./pkg/...; true

lint:
	golint ./pkg/...

clean:
	rm -rf service
	rm -rf manager
	go clean -i .

fmt:
	# go install mvdan.cc/gofumpt@latest
	gofumpt -l -w ./pkg/
	# go fmt ./pkg/...
	# find ./pkg -name "*.go" | xargs gofmt -w

gen-model:
	# go install github.com/xxjwxc/gormt@master
	@gormt

git-push:
	./cmd.sh git push

git-clear:
	./cmd.sh git clear

install:
	@docker plugin install  grafana/loki-docker-driver:latest --alias loki --grant-all-permissions

help:
	@echo "make: compile packages and dependencies"
	@echo "make vet: run specified go vet"
	@echo "make lint: golint ./..."
	@echo "make fmt: gofumpt -l -w ./pkg/"
	@echo "make clean: remove object files and cached files"

