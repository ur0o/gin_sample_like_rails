FROM golang:1.20.5-alpine3.18

RUN apk --update-cache --no-cache add \
    git

ENV WORKDIR /go/src
WORKDIR $WORKDIR

ENTRYPOINT ["sh", "start.sh"]
