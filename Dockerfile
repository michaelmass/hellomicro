FROM alpine:3.9 as builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates
RUN adduser -D -g '' appuser

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

USER appuser

COPY --from=builder /app/app /usr/bin/app

ENTRYPOINT app
