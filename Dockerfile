# Build stage
FROM golang:1.15-alpine3.13 AS build-env
WORKDIR /go/src/github.com/dfinance/dstation
RUN apk add --no-cache make bash
COPY . .

RUN make install LEDGER_ENABLED=false

# Run stage
FROM alpine:3.13

RUN apk --no-cache add ca-certificates
WORKDIR /opt/app
COPY --from=build-env /go/bin/dstation .
