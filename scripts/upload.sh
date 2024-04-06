#!/usr/bin/env bash

set -e
cd "$(dirname "$0")/.."
pwd

TARGET=$1
if [ -z "$TARGET" ]; then
    printf "Syntax: \n\t$0 user@addr\n"
    exit 1
fi

echo "Compiling application..."
./scripts/openapi-http.sh
GOOS=linux GOARCH=amd64 go build -o app -v -ldflags="-w -s" cmd/app/main.go

echo "Enter password several times if asked, that's ok."

ssh -t "$TARGET" "sudo rm /opt/elder/app"

scp -r ./app ./api ./migrations ./web ./scripts/update.sh "$TARGET:/opt/elder"

ssh -t "$TARGET" "cd /opt/elder &&
    (sudo tmux kill-session -t elder-app || pkill tmux);
    chmod +x ./update.sh && sudo ./update.sh &&
    sudo tmux new-session -s elder-app -d \"sudo /opt/elder/app -config-path /opt/elder/config/config.yaml \""
