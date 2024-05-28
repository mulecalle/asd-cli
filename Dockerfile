FROM golang:1.20.5-buster@sha256:eb3f9ac805435c1b2c965d63ce460988e1000058e1f67881324746362baf9572 AS base

# Set destination for COPY
WORKDIR /app
COPY . .

# Build
FROM base as build
RUN export CGO_ENABLED=0 && \
    export GOOS=darwin GOARCH=amd64  && \
    go build -o /asd

# Copy the final binary
FROM scratch AS binary
COPY --from=build /asd /asd
