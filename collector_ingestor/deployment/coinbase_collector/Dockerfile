FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Verify that the dependencies are available
RUN go mod verify

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ARG KAFKA_BOOTSTRAP_SERVERS
ARG SCHEMA_REGISTRY_URL
ENV KAFKA_BOOTSTRAP_SERVERS=$KAFKA_BOOTSTRAP_SERVERS
ENV SCHEMA_REGISTRY_URL=$SCHEMA_REGISTRY_URL

# Write the Kafka bootstrap servers and schema registry URL to the environment file
RUN echo "KAFKA_BOOTSTRAP_SERVERS=$KAFKA_BOOTSTRAP_SERVERS" >> ./collector_ingestor/.env
RUN echo "SCHEMA_REGISTRY_URL=$SCHEMA_REGISTRY_URL" >> ./collector_ingestor/.env

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /usr/local/bin/app ./collector_ingestor/cmd/ingestor/main.go

FROM alpine:3.14

RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

CMD ["app"]
