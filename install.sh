#!/bin/sh
CONFIG_PATH="$HOME/.config/ricer/config.yaml"
EXEC_PATH="/usr/local/bin/"
EXEC_NAME="ricer"

mkdir -p ~/.config/$EXEC_NAME
if [ ! -f "$CONFIG_PATH" ]; then
    cp ./config.yaml.default "$CONFIG_PATH"
fi

go build -o $EXEC_NAME cmd/$EXEC_NAME/main.go
sudo rm -f "$EXEC_PATH$EXEC_NAME"
sudo mv "$EXEC_NAME" "$EXEC_PATH$EXEC_NAME"
