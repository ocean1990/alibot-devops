# disable all default rules
MAKEFLAGS+=-r

.PHONY: all clean build
.DEFAULT: all

all: clean build

clean:
	@echo cleaning...
	@rm -rf *.go
	@rm -rf *.py
	@rm -rf *.html

doc:
	@echo build documents...
	# go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	protoc --doc_out=. --doc_opt=html,index.html *.proto

go:
	@echo build go files
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative *.proto

python:
	@echo build python files
	python3 -m grpc_tools.protoc --python_out=. --grpc_python_out=. --proto_path=. ./*.proto

build:doc go python
