# Building docker file using Multi-stage builds

###########################################################
# STEP-1: Generate an image for creating executable binary
###########################################################

# Start stage, add tag as builder
FROM golang:1.18.3-alpine3.16 AS builder

# Set working directory
WORKDIR $GOPATH/course-management-app

# Copy entire current project directory to working directory in container
COPY . .

# Build 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/main


##############################################
# STEP-2:  Generating a small build image
##############################################

# Start stage from scratch
FROM scratch

# Copy executable binary to new stage
COPY --from=builder /go/bin/main /go/bin/main

# Run the binary
ENTRYPOINT ["/go/bin/main"]
