#!/bin/bash

mkdir -p ./backend/echo
mkdir -p ./frontend/proto

protoc \
    --go_out=plugins=grpc:./backend/echo \
    --plugin=protoc-gen-ts=./frontend/node_modules/.bin/protoc-gen-ts \
    --js_out=import_style=commonjs,binary:./frontend/proto \
    --ts_out=service=true:./frontend/proto \
    ./echo.proto
