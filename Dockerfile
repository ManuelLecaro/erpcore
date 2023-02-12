# Start from golang base image
FROM golang:1.19 as builder

ENV GO111MODULE=on

# Setup folders
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux go build -a -installsuffix cgo -o main ./cmd/erpcore/

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["./main"]
