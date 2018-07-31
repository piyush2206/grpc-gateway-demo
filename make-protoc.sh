#! /bin/bash

# compile proto files to go extension for grpc server

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I time \
--go_out=plugins=grpc:time \
time/time.proto

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I student \
--go_out=plugins=grpc:student \
student/student.proto

#############################################################################
# compile proto files to go gateway extension for grpc reverse proxy server

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I time \
--grpc-gateway_out=logtostderr=true:time \
time/time.proto

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I student \
--grpc-gateway_out=logtostderr=true:student \
student/student.proto