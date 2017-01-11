FROM alpine:3.5

LABEL ImageBaseName=graphql-sample

RUN apk add --update \
    build-base \
    go \
    && rm -rf /var/cache/apk/*

ENV GOPATH=/go
ENV PATH=$PATH:/go/bin

ADD . /go/src/github.com/dbinnersley/graphql-sample

WORKDIR /go/src/github.com/dbinnersley/graphql-sample