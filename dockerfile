FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

WORKDIR /workspaces

COPY ./sample_api .

WORKDIR /workspaces/sample_project/sample_api
RUN go mod download
RUN go mod verify
RUN apt update

