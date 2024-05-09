FROM golang:1.22.0-alpine3.19 AS builder

WORKDIR /usr/local/go/src/

ADD ./ /usr/local/go/src/

RUN go clean --modcache
RUN go build -mod=readonly -o app cmd/main/main.go

FROM alpine:3.19

COPY --from=builder /usr/local/go/src/app /
COPY --from=builder /usr/local/go/src/config.yml /

CMD ["/app"]