#include .env

PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOPATH="$(GOBASE)/vendor:$(GOBASE)"
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID=/tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent


init:
	go get -u github.com/fitan/genapi/tree/main/cmd/genapi
	go get -u github.com/cosmtrek/air
	go get -u github.com/swaggo/swag/cmd/swag

build:
	go build -ldflags "-X main.GitCommitId=`git rev-parse HEAD` -X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'"  -mod vendor -o output/zsk main.go

tidy:
	go mod tidy

watch:
	touch .air.toml && air

gen-ent:
	genapi ent

gen-api:
	genapi router

gen-swag:
	swag init

gen: gen-ent gen-api gen-swag


