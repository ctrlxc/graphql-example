FROM golang:1.15-alpine

WORKDIR /go

RUN apk update \
  && apk add --no-cache git alpine-sdk\
  && go get -u github.com/99designs/gqlgen \
  && go get -u golang.org/x/tools/cmd/goimports \
  && go get -u github.com/lib/pq \
  && go get -u github.com/volatiletech/sqlboiler \
  && go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql \
  && go get -u github.com/graph-gophers/dataloader \
  && go get -u github.com/cosmtrek/air \
  && go get -u github.com/derekparker/delve/cmd/dlv
