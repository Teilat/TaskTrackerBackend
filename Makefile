PROJECTNAME=$(shell basename "$(PWD)")

# Go переменные.
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE):/  #Вы можете удалить или изменить путь после двоеточия.
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Перенаправление вывода ошибок в файл, чтобы мы показывать его в режиме разработки.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID-файл будет хранить идентификатор процесса, когда он работает в режиме разработки
PID=/tmp/.$(PROJECTNAME)-api-server.pid

# Make пишет работу в консоль Linux. Сделаем его silent.
MAKEFLAGS += --silent

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

clean:
	@echo "  >  Cleaning build cache"


vendor:
	@echo "  >  Updating dependencies"
    @GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod tidy
    @GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod vendor
