FROM golang:1.24-alpine3.19 AS build
# Build the catalog service binary with vendored dependencies
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/GooferByte/go-gRPC-MicroSvc
COPY go.mod go.sum ./
COPY vendor vendor
COPY account account
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./catalog/cmd/catalog

FROM alpine:3.11
# Minimal runtime image for the compiled binary
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD [ "app" ]
