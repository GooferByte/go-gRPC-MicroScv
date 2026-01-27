FROM golang:1.25-alpine AS build

# Build the account service binary
RUN apk --no-cache add build-base ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go/bin/app ./account/cmd/account

FROM alpine:3.19
# Minimal runtime image for the compiled binary
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD [ "app" ]
