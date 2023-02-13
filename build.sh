cd /code

export GOPROXY=https://proxy.golang.com.cn,direct
export GOSUMDB=off

go build -o spoved -tags product .
