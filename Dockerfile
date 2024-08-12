FROM golang:1.22-alpine as builder

WORKDIR /home
COPY . .
RUN go build -o /home/helmdocs ./cmd/cmd.go

FROM alpine:latest
COPY --from=builder  /home/helmdocs /opt/helmdocs
ENTRYPOINT ["/opt/helmdocs"]
