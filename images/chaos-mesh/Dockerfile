FROM alpine:3.17

ARG HTTPS_PROXY
ARG HTTP_PROXY

# Upgrade OS packages so alpine:3.17 picks up fixed musl/openssl builds
# (libcrypto3/libssl3 >= 3.0.19-r0, musl/musl-utils >= 1.2.3-r6).
RUN apk upgrade --no-cache && apk add --no-cache tzdata

COPY bin/chaos-controller-manager /usr/local/bin/chaos-controller-manager