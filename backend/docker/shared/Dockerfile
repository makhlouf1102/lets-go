FROM debian:bullseye

RUN apt-get update && apt-get install -y \
    build-essential \
    python3 \
    nodejs \
    gcc \
    g++ \
    git \
    curl

# Install Go (Golang)
ENV GOLANG_VERSION=1.20
RUN curl -fsSL https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"

# Verify Go installation
RUN go version
