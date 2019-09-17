FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/consul-manager

COPY . .

ENV dockerMode on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consul-manager loader.go main.go

FROM alpine:3.9 as prod

WORKDIR /root/consul-manager

COPY --from=0 /go/src/consul-manager  .

CMD ["./consul-manager"]