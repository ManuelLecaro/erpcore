FROM golang:1.18 as builder

ENV GO111MODULE=on

WORKDIR /erp
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go build -gcflags "all=-N -l" -o cmd/erpcore/main .

FROM debian:buster
COPY --from=builder /go/bin/dlv /
COPY --from=builder /erp .
CMD ["./dlv", "--listen=:40000", "--headless=true", "exec", "./main"]
