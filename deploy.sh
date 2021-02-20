#!/bin/bash
git pull
pkill go
pkill go-build
pkill main
go run main.go &
