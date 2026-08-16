FROM golang:1.26.5-alpine AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/atproxy ./cmd/atproxy \
	&& mkdir -p /out/config

FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
	&& adduser -D -H -u 65532 atproxy \
	&& mkdir -p /etc/atproxy \
	&& chown atproxy:atproxy /etc/atproxy

COPY --from=builder /out/atproxy /usr/local/bin/atproxy

USER atproxy

ENV ATPROXY_CONFIG_PATH=/etc/atproxy/config.json

EXPOSE 11111

# atproxy only handles SIGINT (not SIGTERM)
STOPSIGNAL SIGINT

ENTRYPOINT ["atproxy"]
