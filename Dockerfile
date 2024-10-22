FROM golang:alpine3.18 AS builder
RUN apk add --no-cache --update git build-base

WORKDIR /app
COPY . .
RUN go build \
    -a \
    -trimpath \
    -o tele_admin \
    -ldflags "-s -w -buildid=" \
    "./cmd/tele_admin" && \
    ls -lah

FROM alpine:3.18
RUN apk --no-cache add ca-certificates tzdata
ENV LANG=C.UTF-8
ENV LANGUAGE=en_US:en
ENV LC_ALL=C.UTF-8
ENV TZ=UTC
WORKDIR /app

COPY --from=builder /app/tele_admin .
COPY --from=builder /app/conf /app/conf
VOLUME /app/conf
VOLUME /app/log

ENTRYPOINT ["./tele_admin"]