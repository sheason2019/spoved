#!/bin/bash
CURRENT_DIR=$(dirname $0)

cd $CURRENT_DIR

export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

OUTPUT_PATH=$CURRENT_DIR/dist/spoved
SPOVED_PATH=$CURRENT_DIR/cmd/spoved
INITIAL_PATH=$CURRENT_DIR/cmd/initial
SPOVED_INGRESS_PATH=$CURRENT_DIR/cmd/ingress

echo "PRODUCTION = $PRODUCTION"
echo "BUILD_TYPE = $BUILD_TYPE"

buildSpoved() {
  if [ "$PRODUCTION"z = "true"z ]; then
    go build -o $OUTPUT_PATH -tags product $SPOVED_PATH
  else
    go build -o $OUTPUT_PATH $SPOVED_PATH
  fi
}

buildInitial() {
  if [ "$PRODUCTION"z = "true"z ]; then
    go build -o $OUTPUT_PATH -tags product $INITIAL_PATH
  else
    go build -o $OUTPUT_PATH $INITIAL_PATH
  fi
}

buildSpovedIngress() {
  if [ "$PRODUCTION"z = "true"z ]; then
    go build -o $OUTPUT_PATH -tags product $SPOVED_INGRESS_PATH
  else
    go build -o $OUTPUT_PATH $SPOVED_INGRESS_PATH
  fi
}

if [ "$BUILD_TYPE"z = "INITIAL"z ]; then
  buildInitial
elif [ "$BUILD_TYPE"z = "SPOVED"z ]; then
  buildSpoved
elif [ "$BUILD_TYPE"z = "SPOVED_INGRESS"z ]; then
  buildSpovedIngress
fi
