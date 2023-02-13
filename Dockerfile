FROM golang:1.20.0-alpine3.17

WORKDIR /app

COPY ./spoved ./spoved

ENTRYPOINT [ "/app/spoved" ]
