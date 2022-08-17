FROM ubuntu:20.04

RUN apt update && apt install -y arp-scan && apt clean \
    && mkdir /data

COPY templates /app/templates
COPY watchyourlan /app/

EXPOSE 8840

WORKDIR /app

ENTRYPOINT ["./watchyourlan"]