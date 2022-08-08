PROJECT_NAME := $(shell $(basename $(PWD) ))
#
BINARY_NAME := $(shell git config --get remote.origin.url | awk -F/ '{print $$5}' | awk -F. '{print $$1}')
#WIN_BINARY_NAME := $(BINARY_NAME).exe
#BINARY_VERSION := $(shell git describe --tags)
#BINARY_BUILD_DATE := $(shell date +%d.%m.%Y)
#BUILD_FOLDER := .build
#LINTERS = -E asciicheck -E dogsled -E durationcheck -E exportloopref -E forcetypeassert -E goconst -E gocritic -E ifshort -E nilerr -E unconvert -E unparam -E whitespace
## disabled LINTERS: -E tagliatelle
#
ABS_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
ifeq ($(shell go env GOHOSTOS), windows)
	ABS_PATH = $(PWD)
	ifneq ($(shell whereis cygpath), "cygpath:")
		ABS_PATH = $(shell cygpath -w $(CURDIR))
	endif
endif

help:
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help

clean:
	@echo "==> Cleaning build cache"
	rm -r ./vendor
.PHONY: clean

path:
	@echo $(ABS_PATH)
.PHONY: path

vendor:
	@echo "==> Updating dependencies"
	@go mod tidy
	@go mod vendor
.PHONY: vendor

gen-webapi: ## Generate API
	"$(ABS_PATH)/server/api/v1/web/proto":/app $protoc -I=proto --go_out=. --remove_omitempty classes.web.proto
	"$(ABS_PATH)/server/api/v1/web/proto":/app $protoc -I=proto --go_out=. --remove_omitempty elements.web.proto
	"$(ABS_PATH)/server/api/v1/web/proto":/app $protoc -I=proto --go_out=. --remove_omitempty reference_book.web.proto
	"$(ABS_PATH)/server/api/v1/web/proto":/app $protoc -I=proto --web_out=. web.proto


gen:
	protoc -I/usr/local/include -I. \
    	-I${GOPATH}/src \
    	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
    	--grpc-gateway_out=logtostderr=true:./api_pb \
    	--swagger_out=allow_merge=true,merge_file_name=api:. \
    	--go_out=plugins=grpc:./api_pb ./*.proto