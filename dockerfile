FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

WORKDIR /workspaces/sample_api

COPY ./sample_api/go.mod ./sample_api/go.sum ./

# WORKDIR /workspaces/sample_api
RUN go mod download && \
    go mod verify && \
    apt-get update && \
    apt-get install -y vim

