# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.15 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=1 GOOS=linux \
    go build -tags=jsoniter -mod=readonly -v -o thunderstorm \
        -ldflags="-X 'main.Commit=TODO'" \
        .

# Use the official Debian image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim

RUN apt-get update && \
    apt-get install -y ca-certificates  && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/thunderstorm /thunderstorm

# Run the web service on container startup.
CMD ["/thunderstorm"]