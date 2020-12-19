GRPC_GATEWAY_DIR := $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway 2> /dev/null) 
GO_PATH := ${GOPATH}
GO_INSTALLED := $(shell which go)
PROTOC_INSTALLED := $(shell which protoc)
BINDATA_INSTALLED := $(shell which go-bindata 2> /dev/null)
PGGG_INSTALLED := $(shell which protoc-gen-grpc-gateway 2> /dev/null)
PGG_INSTALLED := $(shell which protoc-gen-go 2> /dev/null)
# github.com/gofrs/uuid

all: build

install-tools:
ifndef PROTOC_INSTALLED
	$(error "go is not installed, please run 'brew install go'")
endif
ifndef PROTOC_INSTALLED
	$(error "protoc is not installed, please run 'brew install protobuf'")
endif
ifndef BINDATA_INSTALLED
	@go get -u github.com/kevinburke/go-bindata/go-bindata
endif
ifndef PGGG_INSTALLED
	@go get -u github.com/grpc-ecosystem/grpc-gateway/...
endif
ifndef PGG_INSTALLED
	@go get -u github.com/golang/protobuf/protoc-gen-go
endif

generate: install-tools
	@rm -rf rewardsrefund
	@rm -f www/swagger.json
	@mkdir -p rewardsrefund
	@protoc \
		-I/usr/local/include \
		-I. \
		-I$(GO_PATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:rewardsrefund \
		--swagger_out=logtostderr=true:rewardsrefund \
		--grpc-gateway_out=logtostderr=true:rewardsrefund \
		--proto_path proto rewardsrefund.proto
	@cp rewardsrefund/rewardsrefund.swagger.json www/swagger.json

build: generate
	@rm -rf bin
	@mkdir -p bin
	@go generate ./...
	@go build -o bin/server server/*.go
	@go build -o bin/client client/*.go
	@CGO_ENABLED=0 go build -o bin/server server/*.go
	@echo "Success! Binaries can be found in 'bin' dir"

image:
	@docker build -t grpc-docker .
