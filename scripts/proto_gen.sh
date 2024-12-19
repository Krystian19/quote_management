#!/bin/sh

# Build the clerk's proto gen image
docker build -t quote/proto_gen ./docker/proto_gen

# Generate all the proto defintions found on the /proto directory
docker run -ti -v ${PWD}/proto:/go/src/app/proto -v ${PWD}/makefile:/go/src/app/makefile quote/proto_gen make proto-gen
