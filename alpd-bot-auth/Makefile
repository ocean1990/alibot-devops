# disable all default rules
.SUFFIXES:
MAKEFLAGS+=-r

ifeq (,$(DOCKER_IMAGE))
DOCKER_IMAGE:="alpd-bot-auth:latest"
endif

.PHONY: all clean build test
.DEFAULT: all

all: clean build test

clean:
	@echo cleaning...

test:
	@echo unit testing...
	npm test

build:
	@echo building...
	# build docker image
	docker build -t ${DOCKER_IMAGE} .
