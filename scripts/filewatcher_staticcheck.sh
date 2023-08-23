#!/bin/bash

# To be used with Goland's file watcher

# GoLand file watcher configuration:
# Name: filewatcher_staticcheck
# File type: Go files
# Scope: Project files
# Program: $ProjectFileDir$/scripts/filewatcher_staticcheck.sh
# Arguments: $FileDir$
# In advanced arguments, disable "Auto-save edited files"

PKG_CONFIG_PATH="/opt/homebrew/opt/openssl@3/lib/pkgconfig"

#echo $1
#staticcheck $1
staticcheck -tags dynamic "${1}/..."
