.PHONY: glide deps

glide:
	go get -v github.com/Masterminds/glide
	cd $GOPATH/src/github.com/Masterminds/glide && go install && cd -

deps:
	glide install
