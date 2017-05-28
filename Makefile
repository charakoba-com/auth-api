.PHONY: glide deps

glide:
	mkdir ${GOPATH}/bin
	curl https://glide.sh/get | sh

deps:
	glide install
