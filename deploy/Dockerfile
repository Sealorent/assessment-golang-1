# Build Stage
FROM golang:1.22.0 as build-env

# Set environment variables
ENV APP_NAME mygram
ENV CMD_PATH main.go

# Copy application data into the image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Build the Go application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $CMD_PATH

# Run Stage
FROM alpine:3.14

# Set environment variables
ENV APP_NAME mygram

# Set up the working directory
WORKDIR /app

# Copy SSL certificates from the build stage
COPY --from=build-env /etc/ssl/certs /etc/ssl/certs

# Copy the built executable from the build stage
COPY --from=build-env /$APP_NAME /app/

# Expose application port
EXPOSE 9091

# Start the application
CMD ./$APP_NAME
