FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/service-auth-client

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service-auth-client .


FROM alpine:3.9 as prod

WORKDIR /root/service-auth-client

COPY --from=0 /go/src/service-auth-client  .

CMD ["./service-auth-client"]