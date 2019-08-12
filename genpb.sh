#!/bin/sh

# Generate golang code from *.proto files.
protoc -I pb/ --go_out=plugins=grpc:pb pb/*.proto

