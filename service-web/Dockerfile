FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/WebService

COPY .. .

ENV dockerMode on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o WebService .

FROM alpine:3.9 as prod

WORKDIR /root/WebService

COPY --from=0 /go/src/WebService .

CMD ["./userwebPRJ"]