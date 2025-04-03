#!/bin/bash

PROTO_NAMES=(
    "auth"
    "conversation"
    "errinfo"
    "friend"
    "group"
    "msg"
    "msggateway"
    "push"
    "rtc"
    "sdkws"
    "third"
    "user"
    "statistics"
    "wrapperspb"
)


PLUGIN_PATH="/Users/shengqiang/Documents/Codes/futrtalk-web-embed/node_modules/@protobuf-ts/plugin/bin/protoc-gen-ts"
# shellcheck disable=SC2088
OUTPUT_DIR="/Users/shengqiang/Documents/Codes/futrtalk-web-embed/src/utils/open-im-sdk/model"

rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

if [ ! -d "$OUTPUT_DIR" ]; then
    mkdir -p "$OUTPUT_DIR"
fi

for name in "${PROTO_NAMES[@]}"; do
  ./protoc/protoc \
    --plugin=protoc-gen-ts="$PLUGIN_PATH" \
    --ts_out="$OUTPUT_DIR/" \
    --proto_path=./ \
    ./${name}/${name}.proto

  if [ $? -ne 0 ]; then
      echo "Error processing ${name}.proto"
      exit $?
  fi
done

echo "TypeScript files generated successfully in $OUTPUT_DIR"
