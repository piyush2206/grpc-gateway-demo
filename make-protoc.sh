#! /bin/bash

# compile proto files to go extension for grpc server

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I api/time \
--go_out=plugins=grpc:api/time \
api/time/time.proto

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I api/calc \
--go_out=plugins=grpc:api/calc \
api/calc/calc.proto

#############################################################################
# compile proto files to go gateway extension for grpc reverse proxy server

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I api/time \
--grpc-gateway_out=logtostderr=true:api/time \
api/time/time.proto

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I api/calc \
--grpc-gateway_out=logtostderr=true:api/calc \
api/calc/calc.proto