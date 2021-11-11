FROM golang:latest

ENV LINT_VERSION="v1.42.0"

COPY entrypoint.sh /entrypoint.sh

RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin ${LINT_VERSION} \
  && chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
