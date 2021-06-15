# Build stage
FROM golang:1.15-alpine3.13 AS build-env
WORKDIR /go/src/github.com/dfinance/dstation
RUN apk add --no-cache make bash git
COPY . .

RUN make install LEDGER_ENABLED=false

# Run stage
FROM alpine:3.13

ARG CI_PIPELINE_ID="unset"
ARG CI_COMMIT_REF_NAME="unset"
ARG CI_COMMIT_SHA="unset"

RUN apk --no-cache add ca-certificates
WORKDIR /opt/app
COPY --from=build-env /go/bin/dstation .
