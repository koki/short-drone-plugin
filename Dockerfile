FROM golang:1.8-alpine

RUN apk update && \
    apk add ca-certificates wget && \
    update-ca-certificates

ADD . /go/src/github.com/kubeciio/koki

WORKDIR /go/src/github.com/kubeciio/koki

RUN go build github.com/kubeciio/koki

RUN wget https://github.com/koki/short/releases/download/v0.3.0/short_linux_amd64 && \
    mv short_linux_amd64 /go/bin/short

ENTRYPOINT ["/go/bin/koki"]
