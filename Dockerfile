###########################
# Build executable binary #
###########################

FROM golang:1.16-alpine AS builder

# Add Maintainer Info
LABEL maintainer="Nhanh Dam <damvannhanh48@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy
RUN go mod download

# Copy the sources from the current directory to the Working Directory inside the container
COPY . .

# compile our application
RUN CGO_ENABLED=0 go build -o theworlddata cmd/main.go

###############
# Build image #
###############

FROM gcr.io/distroless/base-debian10

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/theworlddata .

# Run the executable
CMD ["/theworlddata"]