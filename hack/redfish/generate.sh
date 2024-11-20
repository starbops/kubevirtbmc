#!/usr/bin/env sh
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/../../" && pwd)"
cd "$ROOT_DIR"

echo "Working directory: $(pwd)"

command -v openapi-generator || {
  echo "openapi-generator not found. Please install it first."
  exit 1
}

# goimports is used to post-process the generated code
command -v goimports || {
  echo "goimports not found. Please install it first."
  exit 1
}

if [ "$(uname)" = "Darwin" ]; then
    export JAVA_HOME=`/usr/libexec/java_home -v 1.11`
fi

# Generate the models and server stubs from the OpenAPI spec
GO_POST_PROCESS_FILE="goimports -w" openapi-generator generate \
    -i ./hack/redfish/spec/openapi.yaml \
    -o ./pkg/generated/redfish/ \
    -g go-server \
    --package-name server \
    --enable-post-process-file \
    -p sourceFolder=server,onlyInterfaces=true,outputAsLibrary=true,enumClassPrefix=true

