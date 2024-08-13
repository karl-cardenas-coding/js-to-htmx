FROM golang:1.22.6-alpine3.20 AS builder

ARG VERSION

RUN apk update && \
    apk add --no-cache ca-certificates tzdata xdg-utils && \
    update-ca-certificates && \
    adduser -H -u 1002 -D appuser appuser 




ADD ./ /source
WORKDIR /source
RUN go build -o jstohtmx -v 

FROM scratch

ENV PORT=8080

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /source/jstohtmx /jstohtmx

USER appuser

CMD ["/jstohtmx"]