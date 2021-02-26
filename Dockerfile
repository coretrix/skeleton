FROM golang:1.15-buster as builder

ARG APP=web-api

WORKDIR /app

COPY . ./

RUN go build -mod=readonly -v -o main ./cmd/${APP}

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    --no-install-recommends \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main /app/bin/main

COPY ./config /app/config/
COPY ./templates /app/templates/

ENV APP_CONFIG_FOLDER /app/config
ENV APP_TEMPLATE_FOLDER /app/templates

EXPOSE 4001

WORKDIR /app

CMD /app/bin/main
