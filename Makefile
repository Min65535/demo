# note: call scripts from /scripts
GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	mkdir -p api/my/v1/my && rm -rf api/my/v1/my/*
	protoc --proto_path=./api/my/v1 \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api/my/v1/my \
 	       --go-http_out=paths=source_relative:./api/my/v1/my \
 	       --go-grpc_out=paths=source_relative:./api/my/v1/my \
	       $(API_PROTO_FILES)

.PHONY: swagger
# generate swagger file
swagger:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --openapiv2_out . \
		   --openapiv2_opt logtostderr=true \
		   --openapiv2_opt json_names_for_fields=false \
		   $(API_PROTO_FILES)

.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: wire
# generate wire code
wire:
	make wire-outer

.PHONY: wire-router
# generate wire code
wire-router:
	go mod tidy
	cd cmd/router && wire

.PHONY: build
# build
build:
	make wire
	mkdir -p bin/ && go build -ldflags "-s -w -X main.Version=$(VERSION)" -o ./bin/ ./...

# project name
PROJECTNAME=$(shell basename "$(PWD)")

# project path
ROOT=$(shell pwd)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## mysql-init: init mysql db, example.. make mysql-init env=test
mysql-init:
	@echo "init db..."
	@./scripts/init_db_env.sh $(ROOT) $(env)

## test: run project all unit tests
test:
	@echo "run project unit test"
	@./scripts/init_db_env.sh $(ROOT)
	@go test -p 1 -cover ./...

### build: build project, all binaries for programs to $GOPATH/bin
#build:
#	@echo "build project"
#	#@./scripts/shell.sh
#
### mod: update or clear mod pkg, do=tidy  or do=vendor
#mod:
#	@echo "use mod"
#	@./scripts/pkg.sh $(ROOT) $(do)

## docker-image: make docker image, env=test/dev/prod
docker-image:
	@$(ROOT)/gw/docker/mk_image.sh $(env) 1

## docker-image-timer: make docker image, env=test/dev/prod
docker-image-timer:
	@$(ROOT)/gw-timer/docker/mk_image.sh $(env) 1