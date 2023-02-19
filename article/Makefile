GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
GOVERSION:=$(shell go env GOVERSION)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	CONFIG_PROTO_FILES=$(shell $(Git_Bash) -c "find config -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	CONFIG_PROTO_FILES=$(shell find config -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

ifneq ($(wildcard .env.yaml), .env.yaml)
	ENVFILE=$(shell cp .env.example.yaml .env.yaml)
	ENVCREATE_SUCCESS_TIP='创建配置文件成功.'
endif

.PHONY: init
# 初始化安装脚本
init:
	$(ENVFILE)
	@echo $(ENVCREATE_SUCCESS_TIP)
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/google/wire/cmd/wire@latest
	@echo '初始化操作完成!'

.PHONY: generate
# 自动化生成编译前的类库代码
generate:
	go mod download && go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: wire
# 生成依赖注入文件
wire:
	cd ./cmd && $(GOPATH)/bin/wire

.PHONY: config
# 生成配置相关 Proto 文件
config:
	protoc --proto_path=./config \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./config \
	       $(CONFIG_PROTO_FILES)

.PHONY: api
# 生成API相关 Proto 文件
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:./api \
		   --go-http_out=paths=source_relative:./api \
		   --go-grpc_out=paths=source_relative:./api \
           --validate_out=paths=source_relative,lang=go:./api \
	   	   --go-errors_out=paths=source_relative:./api \
		   --openapi_out=fq_schema_naming=true,default_response=false:. \
		   --openapiv2_out ./api \
	   	   --openapiv2_opt logtostderr=true \
	   	   --openapiv2_opt json_names_for_fields=false \
	       $(API_PROTO_FILES)

.PHONY: gormgen
# 生成数据库查询器文件
gormgen:
	cd generate/gormgen && go run main.go

.PHONY: run
# 开发环境启动项目
run:
	cd cmd && go run ./...

.PHONY: build
# 编译构建项目 (需要有 .git 提交版本)
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(GIT_VERSION)" -o ./bin/server ./cmd

# 帮助命令
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
