#!/usr/bin/env bash

export GOPATH=/root/go
cd ~/go/src/github.com/bom-d-van/me

echo "Update Me"
git pull --rebase

echo "DONE"
