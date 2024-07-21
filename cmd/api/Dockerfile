# The main Build image to build our binaries
FROM golang:1.22.3-alpine3.19 as build

# Install UPX for binary compression
RUN apk add upx

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code of the component
COPY ./cmd/api/api.go ./cmd/api/api.go

# build binary component
RUN go build -ldflags="-s -w" -o api.bin ./cmd/api/api.go

# Create release image
FROM alpine:3.19 as release

# Copy our static executable
COPY --from=build /app/api.bin /api

EXPOSE 8080
CMD [ "/api" ]