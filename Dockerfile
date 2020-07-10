FROM alpine:latest

WORKDIR /opt/simple-api

EXPOSE 8080

ADD ./target/simple-api simple-api

ENTRYPOINT ["/opt/simple-api/simple-api"]
