FROM golang:alpine AS builder

RUN apk add build-base

WORKDIR /src
COPY go.mod go.sum .
RUN go mod download

COPY . .
WORKDIR /src/cmd/WatchYourLAN
RUN CGO_ENABLED=0 go build -o /WatchYourLAN .


FROM alpine

WORKDIR /app

RUN apk add --no-cache arp-scan tzdata \
    && mkdir /data

COPY --from=builder /WatchYourLAN /app/

ENTRYPOINT ["./WatchYourLAN"]