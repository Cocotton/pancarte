FROM golang:1.8

RUN apt-get update && apt-get install -y --no-install-recommends git \
	&& rm -rf /var/lib/apt/lists/*

COPY . $GOPATH/src/github.com/cocotton/pancarte
RUN go get github.com/cocotton/pancarte
RUN go install github.com/cocotton/pancarte

ENV PANCARTE_DB_HOST    someHost
ENV PANCARTE_DB_NAME    pancarte
ENV PANCARTE_PORT       somePort
ENV PANCARTE_JWT_SECRET someSecret

ENTRYPOINT $GOPATH/bin/pancarte
