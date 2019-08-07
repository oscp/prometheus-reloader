FROM golang as builder
WORKDIR /prometheus-reloader
COPY . .
RUN go get -v

FROM centos:7
COPY --from=builder /go/bin/main /usr/local/bin
ENTRYPOINT main
