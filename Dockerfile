FROM golang:1.19.0-buster@sha256:ba38b423b8a7ef6524d71086c654a157561e46d36d0a2e10e8d60c8d4f0763b4 AS base

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
