#!/bin/bash
GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-shared -o GoWallhack.dll
GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-archive