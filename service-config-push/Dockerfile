FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/service-config-push

COPY . .

ENV dockerMode on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service-config-push loader.go main.go

FROM alpine:3.9 as prod

WORKDIR /root/service-config-push

COPY --from=0 /go/src/service-config-push  .

CMD ["./service-config-push"]