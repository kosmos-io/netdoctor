FROM alpine:3.17.1

ARG BINARY

RUN apk add --no-cache ca-certificates
RUN apk update && apk upgrade
RUN apk add --no-cache ip6tables iptables curl netcat-openbsd

COPY ${BINARY} /bin/${BINARY}
