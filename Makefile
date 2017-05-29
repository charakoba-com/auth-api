.PHONY: glide deps initdb

glide:
	mkdir ${GOPATH}/bin
	curl https://glide.sh/get | sh

deps:
	glide install

initdb:
	@echo "initializing database..."
	@mysql -u root -e "DROP DATABASE IF EXISTS authapi;"
	@mysql -u root -e "CREATE DATABASE authapi;"
	@mysql -u root authapi < sql/authapi.sql
