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
  if ! gocyclo -over 15 .; then
    exit 1
  fi
  if ! go test -cover -race -coverprofile=cover.out -outputdir=coverage .; then
    exit 1
  fi
  if ! go tool cover -html=./coverage/cover.out -o ./coverage/cover.out.html; then
    exit 1
  fi
}

COMMAND=$1

case "$COMMAND" in
test    )
  runTests ;;
benchmark    )
  go test -benchmem -cpuprofile=cpu.out -memprofile=mem.out -outputdir=coverage -bench .
  go tool pprof -svg coverage/cpu.out > coverage/cpu.out.svg
  go tool pprof -svg coverage/mem.out > coverage/mem.out.svg
  ;;
*       )
  echo
  echo "Commands"
  echo "test - run tests"
  echo "benchmark - run test benchmarks"
  ;;
esac
