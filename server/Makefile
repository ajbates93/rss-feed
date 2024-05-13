# Include variables from the .envrc file
include .envrc

## help: print this help message
.PHONY: help
help: 
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# DEVELOPMENT --------------------------

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${RSS_FEED_DSN}
