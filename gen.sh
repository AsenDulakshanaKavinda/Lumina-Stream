#!/bin/bash

# Define paths
PROTO_SRC="proto/embedding.proto"
GO_OUT="go-orchestrator/pb"
PY_OUT="python-embedder/pb"

# Ensure Go binaries are in the PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Create directories if they don't exist
mkdir -p $GO_OUT
mkdir -p $PY_OUT

echo "--- Generating Python files ---"
# Ensure we use the venv python if active
python3 -m grpc_tools.protoc -Iproto \
       --python_out=$PY_OUT \
       --grpc_python_out=$PY_OUT \
       --pyi_out=$PY_OUT \
       $PROTO_SRC

# Ensure Python treats the pb directory as a package
touch $PY_OUT/__init__.py

echo "--- Generating Go files ---"
protoc -Iproto \
       --go_out=$GO_OUT --go_opt=paths=source_relative \
       --go_opt=Membedding.proto=./pb \
       --go-grpc_out=$GO_OUT --go-grpc_opt=paths=source_relative \
       --go-grpc_opt=Membedding.proto=./pb \
       $PROTO_SRC