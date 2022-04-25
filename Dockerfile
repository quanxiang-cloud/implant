FROM alpine as certs
RUN apk update && apk add ca-certificates

FROM golang:1.18.1-alpine3.15 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o implant -mod=vendor -ldflags='-s -w'  -installsuffix cgo cmd/main.go

FROM scratch
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

WORKDIR /faas
COPY --from=builder ./build/implant ./cmd/

ENTRYPOINT ["./cmd/implant","-fn-update=$FN_UPDATE", "-doc-update=$DOC_UPDATE"]
