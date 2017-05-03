FROM golang:1.7.5-alpine3.5

RUN apk update && apk add --no-cache git

COPY . $GOPATH/src/github.com/cocotton/pancarte
RUN go get github.com/cocotton/pancarte
RUN go install github.com/cocotton/pancarte

ENV PANCARTE_DB_HOST    someHost
ENV PANCARTE_DB_NAME    pancarte
ENV PANCARTE_PORT       somePort
ENV PANCARTE_JWT_SECRET someSecret

ENTRYPOINT $GOPATH/bin/pancarte
