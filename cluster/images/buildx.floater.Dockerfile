FROM alpine:3.17.1

ARG BINARY
ARG TARGETPLATFORM

RUN apk add --no-cache ca-certificates
RUN apk update && apk upgrade
RUN apk add ip6tables iptables curl netcat-openbsd

COPY ${TARGETPLATFORM}/${BINARY} /bin/${BINARY}
