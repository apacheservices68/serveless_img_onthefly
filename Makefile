OK_COLOR=\033[32;01m
NO_COLOR=\033[0m

export GO111MODULE=auto

build:
	@echo "$(OK_COLOR)==> Compiling binary$(NO_COLOR)"
	go test
	go build -o dist/tto-resize

test:
	go test

install:
	go get -u .

deploy:
	yes | cp -rf ./dist/tto-resize "$(GOPATH)/dist/tto-resize"
	@echo "$(OK_COLOR)==> Deploy sucess$(NO_COLOR)"

benchmark: build
	bash benchmark.sh
