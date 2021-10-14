APP_DIR="/go/src/github.com/${GITHUB_REPOSITORY}/"

mkdir -p "${APP_DIR}" && cp -r ./ "${APP_DIR}" && cd "${APP_DIR}"

export GO111MODULE=on
go mod tidy
go mod verify

if [[ "$1" == "lint" ]]; then
    echo " Running GolangCI-Lint..."
    golangci-lint --version
    echo
    golangci-lint run -c .golangci_config.yml --tests=False --timeout=30m --max-issues-per-linter 0 --max-same-issues 0 --out-format=colored-line-number ./...
fi
