FROM golang:alpine AS builder
WORKDIR /go/src/github.com/filetrust/icap-service-metrics-exporter
COPY . .
RUN cd cmd \
    && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o metrics-exporter .

FROM scratch
COPY --from=builder /go/src/github.com/filetrust/icap-service-metrics-exporter/cmd/metrics-exporter /bin/metrics-exporter

ENTRYPOINT ["/bin/metrics-exporter"]