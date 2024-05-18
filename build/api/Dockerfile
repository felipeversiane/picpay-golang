FROM golang:1.22-alpine

RUN mkdir /app

ENV GIN_MODE=${GIN_MODE:-release}

ADD . /app

WORKDIR /app

RUN go build -o main cmd/server/main.go

CMD ["/app/main"]