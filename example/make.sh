#!/usr/bin/env bash

init() {
  PKG_LIST=$(go list ./... | grep -v /vendor/ | grep -v /lib/)
}

case $1 in

install)
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest
  go install github.com/swaggo/swag@latest
  ;;

lint)
  golangci-lint -c .golangci.yml run
  ;;

test)
  init
  go test -count=1 ${PKG_LIST}
  ;;

run)
  go run ./cmd/server/*.go
  ;;

build)
  go build -o app ./cmd/server/*.go && chmod +x app
  ;;

swag-doc)
  swag init --parseDependency -g cmd/server/*.go -o api
  ;;

help)
  cat make.sh | grep "^[a-z-]*)"
  ;;

*)
  echo "unknown $1, try help"
  ;;

esac
