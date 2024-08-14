FROM alpine:latest
COPY ./helmdocs /opt/helmdocs
ENTRYPOINT ["/opt/helmdocs"]
