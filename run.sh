#!/bin/bash

# simple way to kick off runs of the project, since 'go run' sucks!
make build || exit 1
sudo ./gd3 "$@"
e=$?
make clean
exit $e
