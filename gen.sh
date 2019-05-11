#!/bin/bash

mkdir -p ./server/echo
mkdir -p ./front/proto

protoc \
    --go_out=plugins=grpc:./server/echo \
    --plugin=protoc-gen-ts=./front/node_modules/.bin/protoc-gen-ts \
    --js_out=import_style=commonjs,binary:./front/proto \
    --ts_out=service=true:./front/proto \
    ./echo.proto
