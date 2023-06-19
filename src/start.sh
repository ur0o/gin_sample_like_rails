#! /bin/sh

if [ ! -f ./go.mod ]; then
    go mod init gin_sample
    go get -u github.com/cosmtrek/air
    go build -o /go/bin/air github.com/cosmtrek/air
fi
go mod tidy

air -c air.toml