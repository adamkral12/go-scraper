FROM golang:latest AS builder

ADD . /app
WORKDIR /app
RUN go build -o main

FROM debian:jessie-slim
RUN mkdir /app
COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]