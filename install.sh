#!/bin/sh

go build -o ricer cmd/ricer/main.go
sudo mv ricer /usr/local/bin/
