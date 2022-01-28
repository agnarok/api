FROM golang:1.15

RUN mkdir -p /go/src/github.com/agnarok/api

WORKDIR /go/src/github.com/agnarok/api

COPY . /go/src/github.com/agnarok/api

RUN make setup

RUN go build -v -o api

CMD ["./api"]
