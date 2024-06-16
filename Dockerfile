FROM golang:alpine3.18 as builder
WORKDIR /golang-fiber-in-docker
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" .
FROM busybox
WORKDIR /golang-fiber-in-docker
COPY --from=builder /golang-fiber-in-docker /usr/bin/
ENTRYPOINT ["golang-fiber-in-docker"]