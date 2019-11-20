# note: call scripts from /scripts

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

## test: run prject all unit tests
test:
	@echo "run project unit test"
	@./scripts/init_db_env.sh $(ROOT)
	@go test -p 1 -cover ./...

## build: build project, all binaries for programs to $GOPATH/bin
build:
	@echo "build project"
	#@./scripts/shell.sh

## mod: update or clear mod pkg, do=tidy  or do=vendor
mod:
	@echo "use mod"
	@./scripts/pkg.sh $(ROOT) $(do)

## docker-image: make docker image, env=test/dev/prod
docker-image:
	@$(ROOT)/gw/docker/mk_image.sh $(env) 1

## docker-image-timer: make docker image, env=test/dev/prod
docker-image-timer:
	@$(ROOT)/gw-timer/docker/mk_image.sh $(env) 1