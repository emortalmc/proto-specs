#!/bin/bash

PATH="$PATH:$(go env GOPATH)/bin"
export PATH

rm -rf go/{grpc,message,model}
mkdir -p go

find ../src/proto -name '*.proto' -print0 | while IFS= read -r -d '' file
do
  protoc -I ../src/proto --go_out=go --go-grpc_out=go "$file"
done

mv go/github.com/emortalmc/proto-specs/gen/go/* go/
rm -rf go/github.com
