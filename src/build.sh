#!/bin/bash
work_path=$(cd `dirname $0`/../; pwd)
echo $work_path

GOPATH=$GOPATH:/Users/homandeng/code/serverless-transitcode GOOS=linux GOARCH=amd64 go build -o main *.go
zip main.zip main
