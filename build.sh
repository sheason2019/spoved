CURRENT_DIR=$(dirname $0)

cd $CURRENT_DIR

export GOPROXY=https://proxy.golang.com.cn,direct
export GOSUMDB=off

if [ "$PRODUCT"z = "true"z ]; then
  echo "build production"
  go build -o spoved -tags product ./cmd/spoved
else
  echo "build development"
  go build -o spoved ./cmd/spoved
fi
