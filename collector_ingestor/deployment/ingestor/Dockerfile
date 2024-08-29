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
ARG DB_USER
ARG DB_PASSWORD
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME

ENV KAFKA_BOOTSTRAP_SERVERS=$KAFKA_BOOTSTRAP_SERVERS
ENV SCHEMA_REGISTRY_URL=$SCHEMA_REGISTRY_URL
ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_HOST=$DB_HOST
ENV DB_PORT=$DB_PORT
ENV DB_NAME=$DB_NAME

# Write the Kafka bootstrap servers and schema registry URL to the environment file
RUN echo "KAFKA_BOOTSTRAP_SERVERS=$KAFKA_BOOTSTRAP_SERVERS" >> ./collector_ingestor/.env
RUN echo "SCHEMA_REGISTRY_URL=$SCHEMA_REGISTRY_URL" >> ./collector_ingestor/.env
RUN echo "DB_USER=$DB_USER" >> ./collector_ingestor/.env
RUN echo "DB_PASSWORD=$DB_PASSWORD" >> ./collector_ingestor/.env
RUN echo "DB_HOST=$DB_HOST" >> ./collector_ingestor/.env
RUN echo "DB_PORT=$DB_PORT" >> ./collector_ingestor/.env
RUN echo "DB_NAME=$DB_NAME" >> ./collector_ingestor/.env

# Build the Go app
RUN go build -v -o /usr/local/bin/app ./collector_ingestor/cmd/ingestor/main.go

CMD ["app"]
