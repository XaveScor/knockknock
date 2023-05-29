FROM golang:1.20-alpine3.18

WORKDIR /app
COPY . /app

RUN go build -o main .

FROM alpine:3.14

WORKDIR /app
ENV REDIS_ADDR=synology.local:6379
COPY --from=0 /app/main .

ENTRYPOINT ["./main"]
