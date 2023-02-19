FROM golang:1.20.0-alpine3.17

WORKDIR /app

COPY ./dist/spoved ./dist/spoved

ENTRYPOINT [ "/app/dist/spoved" ]
