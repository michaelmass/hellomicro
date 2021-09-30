FROM golang:1.17-alpine as builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates
RUN adduser -D -g '' appuser

COPY . .
RUN CGO_ENABLED=0 go build .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

USER appuser

COPY --from=builder /app/hellomicro /usr/bin/hellomicro

ENTRYPOINT ["hellomicro", "serve"]
