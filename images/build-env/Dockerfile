# syntax=docker/dockerfile:experimental

FROM debian:buster-slim

ENV DEBIAN_FRONTEND noninteractive

ARG HTTPS_PROXY
ARG HTTP_PROXY

ENV http_proxy $HTTP_PROXY
ENV https_proxy $HTTPS_PROXY

RUN apt-get update && \
    apt-get install build-essential curl git pkg-config libfuse-dev fuse -y && \
    rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs && \
    rm -rf /var/lib/apt/lists/*
RUN npm install pnpm@8 -g

ARG TARGET_PLATFORM=amd64

RUN case "$TARGET_PLATFORM" in \
    'linux/amd64') \
    export GO_ARCH='amd64'; \
    ;; \
    'linux/arm64') \
    export GO_ARCH='arm64'; \
    ;; \
    *) echo >&2 "error: unsupported architecture '$TARGET_PLATFORM'"; exit 1 ;; \
    esac; \
    curl -vvv https://dl.google.com/go/go1.20.8.linux-$GO_ARCH.tar.gz | tar -xz -C /usr/local |

ENV PATH "/usr/local/go/bin:${PATH}"
ENV GOCACHE /tmp/go-build
ENV GOPATH /tmp/go

LABEL org.opencontainers.image.source https://github.com/chaos-mesh/chaos-mesh
