FROM golang:1.22.2-bullseye AS builder
WORKDIR /workspace
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY api/ api/
COPY pkg/ pkg/
COPY main.go ./
COPY Makefile .
RUN make client-build

FROM debian:bullseye-slim
RUN apt-get update && apt-get install ca-certificates iproute2 -y && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder /workspace/bin/inonius_v3cli /bin/inonius_v3cli

ENTRYPOINT ["/bin/inonius_v3cli"]
