.PHONY: glide deps initdb inittest

glide:
	mkdir ${GOPATH}/bin
	curl https://glide.sh/get | sh

deps:
	glide install

initdb:
	@echo "initializing database..."
	@mysql -u root -e "DROP DATABASE IF EXISTS apidb;"
	@mysql -u root -e "CREATE DATABASE apidb;"
	@mysql -u root apidb < sql/authapi.sql

inittest: initdb
	@mysql -u root apidb < sql/test_data.sql
