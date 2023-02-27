#!/bin/bash
CURRENT_DIR=$(dirname $0)

cd $CURRENT_DIR

export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

OUTPUT_PATH=$CURRENT_DIR/dist/spoved
SPOVED_PATH=$CURRENT_DIR/cmd/spoved
INITIAL_PATH=$CURRENT_DIR/cmd/initial

echo "PRODUCT = $PRODUCT"
echo "BUILD_TYPE = $BUILD_TYPE"

buildSpoved() {
  if [ "$PRODUCT"z = "true"z ]; then
    go build -o $OUTPUT_PATH -tags product $SPOVED_PATH
  else
    go build -o $OUTPUT_PATH $SPOVED_PATH
  fi
}

buildInitial() {
  if [ "$PRODUCT"z = "true"z ]; then
    go build -o $OUTPUT_PATH -tags product $INITIAL_PATH
  else
    go build -o $OUTPUT_PATH $INITIAL_PATH
  fi
}

if [ "$BUILD_TYPE"z = "INITIAL"z ]; then
  buildInitial
elif [ "$BUILD_TYPE"z = "SPOVED"z ]; then
  buildSpoved
fi
