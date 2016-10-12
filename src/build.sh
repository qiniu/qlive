#!/usr/bin/env bash
DIR=$(cd ../; pwd)
export GOPATH=$GOPATH:$DIR
go build main.go
