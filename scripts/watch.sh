#!/usr/bin/env bash

find . -name \*.go -print | entr -r go run main.go