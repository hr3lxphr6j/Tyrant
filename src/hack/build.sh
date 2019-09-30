#!/bin/bash

cd ${1:-"./"}
go build -o bin/Tyrant.$(go env GOOS) ./src