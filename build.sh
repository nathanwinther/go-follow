#!/bin/sh -xv

export GOPATH=`pwd`

go get -u github.com/mattn/go-sqlite3
go get -u github.com/nathanwinther/go-feedparser

go install follow

