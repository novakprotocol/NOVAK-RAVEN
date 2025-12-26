#!/bin/bash
set -e
cd /opt/novak-src
git pull
go build -o /opt/novak/bin/novakd ./cmd/novakd
sudo systemctl restart novakd
echo "[NOVAK] updated $(date -u)"
