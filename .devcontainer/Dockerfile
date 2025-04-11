FROM mcr.microsoft.com/devcontainers/base:ubuntu-24.04

# Install golang
# https://go.dev/dl/
ARG GOLANG_VERSION="1.24.1"
ARG GOLANG_CHECKSUM="cb2396bae64183cdccf81a9a6df0aea3bce9511fc21469fb89a0c00470088073"
RUN wget --https-only "https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" -O /tmp/golang.tar.gz && \
    sha256sum /tmp/golang.tar.gz | grep "${GOLANG_CHECKSUM}" && \
    tar -C /usr/local -xf /tmp/golang.tar.gz && \
    rm /tmp/golang.tar.gz
ENV GOROOT="/usr/local/go"
ENV GOPATH="/home/vscode/go"
ENV PATH="${PATH}:${GOROOT}/bin:${GOPATH}/bin"

# Install golang tooling
RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/tools/gopls@latest