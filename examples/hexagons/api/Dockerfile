FROM golang:1.10.1-alpine3.7 as builder
WORKDIR /usr/local/go/src/github.com/dgageot/hexagons/
COPY vendor vendor
RUN find vendor -maxdepth 2 -mindepth 2 -type d -exec sh -c 'go install -i github.com/dgageot/hexagons/{}/... || true' \;
COPY main.go .
RUN go build -o /api main.go

FROM alpine:3.7
EXPOSE 8080
ENTRYPOINT ["/api"]
COPY --from=builder /api /
