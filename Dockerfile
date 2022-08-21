FROM golang:alpine AS builder

RUN apk add build-base
COPY src /src
RUN cd /src && go build .


FROM alpine

RUN apk add arp-scan \
    && mkdir /data

COPY src/templates /app/templates
COPY --from=builder /src/watchyourlan /app/

EXPOSE 8840

WORKDIR /app

ENTRYPOINT ["./watchyourlan"]