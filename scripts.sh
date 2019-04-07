#!/usr/bin/env bash

runTests() {
  if ! go vet .; then
    exit 1
  fi
  if ! golint -set_exit_status .; then
    exit 1
  fi
  if ! gocritic check -enable='#diagnostic,#style,#performance' -disable='docStub' .; then
    exit 1
  fi
  if ! go test -cover -race -coverprofile=app.coverprofile -outputdir=coverage .; then
    exit 1
  fi
  if ! go tool cover -html=./coverage/app.coverprofile -o ./coverage/app.coverprofile.html; then
    exit 1
  fi
}

COMMAND=$1

case "$COMMAND" in
test    )
  runTests ;;
*       )
  echo
  echo "Commands"
  echo "test - run tests"
  ;;
esac
