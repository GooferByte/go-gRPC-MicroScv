FROM golang:1.25-alpine AS build

RUN apk --no-cache add build-base ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/app ./graphql

FROM alpine:3.19

WORKDIR /usr/bin

COPY --from=build /go/bin .

EXPOSE 8000

CMD ["app"]
