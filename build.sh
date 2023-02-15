cd /code


export GOPROXY=https://proxy.golang.com.cn,direct
export GOSUMDB=off

if [ "$PRODUCT"z == "true"z ]; then
  echo "build production"
  go build -o spoved -tags product .
else
  echo "build development"
  go build -o spoved .
fi
