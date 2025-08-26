#!/usr/bin/env nix-shell
#!nix-shell -i bash
#!nix-shell -p htmlq -p go

./build.sh
go run . -h < options.txt > filtered.html
te formatted.html.tet

