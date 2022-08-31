# Start from golang base image
FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
RUN go mod download
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
# RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY ./app .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

ENV GOPATH=/usr/src/app/

RUN curl -OL https://golang.org/dl/go1.17.linux-arm64.tar.gz; mkdir /etc/golang; tar -xvzf go1.17.linux-arm64.tar.gz -C /etc/golang; ln -s /etc/golang/go/bin/go /usr/bin/go; rm -f go1.17.linux-arm64.tar.gz
RUN go get github.com/jsonnet-bundler/jsonnet-bundler/cmd/jb@latest; ln -s /root/go/bin/jb /usr/bin/jb

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["/root/main"]