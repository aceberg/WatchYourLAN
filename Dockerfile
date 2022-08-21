FROM golang:alpine AS builder

RUN apk add build-base
COPY src /src
RUN cd /src && go build .


FROM alpine

RUN apk add --no-cache tzdata arp-scan \
    && mkdir /data

COPY src/templates /app/templates
COPY --from=builder /src/watchyourlan /app/

WORKDIR /app

ENTRYPOINT ["./watchyourlan"]