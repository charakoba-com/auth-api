.PHONY: glide deps initdb inittest

glide:
	os=$(echo `uname` | tr '[:upper:]' '[:lower:]')
	@echo "OS=${os}"
	@echo type "curl"
	@echo type "wget"
	mkdir ${GOPATH}/bin
	curl https://glide.sh/get | sh

deps:
	glide install

initdb:
	@echo "initializing database..."
	@mysql -u root -e "DROP DATABASE IF EXISTS apidb;"
	@mysql -u root -e "CREATE DATABASE apidb;"
	@mysql -u root apidb < sql/authapi.sql

inittest:
	@mysql -u root apidb < sql/test_data.sql
