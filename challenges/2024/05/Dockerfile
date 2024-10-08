FROM debian:bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -y socat && rm -rf /var/lib/apt/lists/*

COPY chall /chall

CMD socat \
      TCP-LISTEN:1337,reuseaddr,fork \
      EXEC:"/chall"
