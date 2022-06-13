SHELL := /bin/bash

run:
	go run cmd/main.go 


tidy:
	go mod tidy


KIND_CLUSTER := fwc-claster

postgres:
	docker run --name basic-postgres --rm -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e PGDATA=/var/lib/postgresql/data/pgdata1 -v /tmp:/var/lib/postgresql/data1 -p 5432:5432 -it postgres:13.1-alpine
