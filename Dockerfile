FROM golang:1.20.0-alpine3.17

WORKDIR /code

COPY . .

RUN sh build.sh

ENTRYPOINT [ "/code/spoved" ]
