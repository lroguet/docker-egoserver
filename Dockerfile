FROM golang:alpine as builder
WORKDIR /app
COPY cmd/egoserver/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o egoserver .

FROM scratch
COPY --from=builder /app/egoserver /egoserver
CMD ["/egoserver"]
