FROM golang:1.21-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bin/ltak-server ./cmd/server

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /bin/ltak-server /usr/local/bin/ltak-server
WORKDIR /srv/ltak
COPY assets ./assets
EXPOSE 8080
VOLUME ["/srv/ltak/assets"]
ENTRYPOINT ["/usr/local/bin/ltak-server"]
