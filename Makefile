.PHONY: glide deps

glide:
	curl https://glide.sh/get | sh

deps:
	glide install
