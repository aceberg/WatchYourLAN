FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/WatchYourLAN/ && CGO_ENABLED=0 go build -o /WatchYourLAN .


FROM alpine

WORKDIR /app

RUN apk add --no-cache arp-scan curl tzdata \
    && mkdir /data

COPY --from=builder /WatchYourLAN /app/

HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost:8840 || exit 1

ENTRYPOINT ["./WatchYourLAN"]