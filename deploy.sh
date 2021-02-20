#!/bin/bash
git pull
pkill go
pkill go-build
go run main.go &
