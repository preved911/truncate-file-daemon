FROM golang:1.16 as builder

ENV CGO_ENABLED=0

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN go build -o truncate cmd/truncate/main.go

FROM scratch
COPY --from=builder /build/truncate /usr/local/bin/truncate
ENTRYPOINT ["/usr/local/bin/truncate"]
