FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY . .
RUN CGO_ENABLED=0 go build .

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/hellomicro /usr/bin/hellomicro

EXPOSE 8080

CMD ["serve"]
ENTRYPOINT ["hellomicro"]
