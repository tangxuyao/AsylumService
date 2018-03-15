#!/usr/bin/env bash

mkdir $GOPATH/src/proto/asylum
protoc --go_out=plugins=micro:$GOPATH/src/proto/asylum ./asylum.api.proto