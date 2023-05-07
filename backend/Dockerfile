FROM golang:1.20-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

FROM debian:buster-slim
ENV IS_DEPLOYMENT=true

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates curl netcat && \
    rm -rf /var/lib/apt/lists/*

RUN curl -L https://release.ariga.io/atlas/atlas-linux-amd64-latest -o /tmp/atlas-latest
RUN install -o root -g root -m 0755 /tmp/atlas-latest /usr/local/bin/atlas

WORKDIR /app

COPY --from=builder /app/server /app/server

COPY ent/migrate/migrations /app/migrations
# should be mounted via volume
COPY public /app/public

COPY migrate.sh /app/migrate.sh
RUN chmod +x /app/migrate.sh

COPY docker-entrypoint.sh /

ENTRYPOINT [ "/bin/bash", "/docker-entrypoint.sh" ]
CMD ["/app/server"]