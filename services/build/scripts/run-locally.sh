#!/usr/bin/bash

dotenv() {
  set -o allexport
  [[ -f .env ]] && source .env
  set +o allexport
}


dotenv && go run "${PWD}/cmd/${1}/main.go"
