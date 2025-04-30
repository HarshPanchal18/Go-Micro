#!/bin/bash

set -e

PROTO_DIR=proto
OUT_DIR=.

echo "🔄 Compiling Protobuf files..."

protoc --go_out=$OUT_DIR --go-grpc_out=$OUT_DIR $PROTO_DIR/*.proto

echo "✅ Proto files generated successfully!"