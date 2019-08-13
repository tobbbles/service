FROM golang:1.12-alpine as builder
USER root
RUN apk add --no-cache git
WORKDIR /service/
COPY . .
ENV GOOS=linux GOARCH=amd64 CGOENABLED=0
RUN go build -o /service/service

FROM alpine:3.10 as application
RUN apk add --no-cache ca-certificates
WORKDIR /opt
COPY --from=builder /service/service /opt/service
EXPOSE 3000
ENTRYPOINT ["/opt/service"]
