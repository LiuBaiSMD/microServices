FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/service-discuss-room

COPY . .

ENV dockerMode on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service-discuss-room main.go plugin.go

FROM alpine:3.9 as prod

WORKDIR /root/service-discuss-room

COPY --from=0 /go/src/service-discuss-room  .

CMD ["./service-discuss-room"]