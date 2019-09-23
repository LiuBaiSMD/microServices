FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/service-web

COPY . .

ENV dockerMode on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service-web main.go plugin.go

FROM alpine:3.9 as prod

WORKDIR /root/service-web

COPY --from=0 /go/src/service-web .

CMD ["./service-web"]