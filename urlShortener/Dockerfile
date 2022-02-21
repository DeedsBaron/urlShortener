FROM golang:1.17 AS builder
COPY . /urlShortener
WORKDIR /urlShortener
RUN make build

FROM debian:buster
ARG SOLUTION
ENV SOLUTION=${SOLUTION}
COPY --from=builder /urlShortener ./
COPY init.sh /
CMD bash init.sh