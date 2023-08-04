FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git make

WORKDIR $GOPATH/src/github.com/Lachann/rrs
COPY . .

RUN CGO_ENABLED=0 make build

FROM scratch

COPY --from=builder /go/src/github.com/Lachann/rrs/dist/rrs /usr/local/bin/rrs

ENTRYPOINT ["/usr/local/bin/rrs"]

EXPOSE 8080


