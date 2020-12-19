FROM golang:1.15-alpine

WORKDIR /go

RUN apk update \
  && apk add --no-cache git alpine-sdk \
  && go get -u github.com/lib/pq \
  && go get -u github.com/99designs/gqlgen \
  && go get -u github.com/volatiletech/sqlboiler \
  && go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql