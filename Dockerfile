FROM golang:1.14-alpine as builder

WORKDIR /app
COPY . .

RUN apk --no-cache add git alpine-sdk build-base gcc

RUN go get github.com/kardianos/service
RUN go get github.com/streadway/amqp

RUN go build -o main cmd/agent/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]
