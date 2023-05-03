FROM golang:latest AS builder

RUN update-ca-certificates

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go install -v ./cmd

FROM scratch

COPY --from=builder /go/bin/cmd /go/bin/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD ["/go/bin/app"]

