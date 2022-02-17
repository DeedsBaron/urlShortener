FROM golang:1.17 AS builder

COPY . /urlShortener

WORKDIR /urlShortener

RUN make build

FROM debian:buster

COPY --from=builder /urlShortener ./

CMD ["./urlShortener"]