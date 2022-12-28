FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/WatchYourLAN/ && CGO_ENABLED=0 go build -o /WatchYourLAN .


FROM alpine

WORKDIR /app

RUN apk add --no-cache tzdata arp-scan \
    && mkdir /data

COPY --from=builder /WatchYourLAN /app/

ENTRYPOINT ["./WatchYourLAN"]