#!/usr/bin/env nix-shell
#!nix-shell -i bash
#!nix-shell -p htmlq -p gnused -p go

htmlq 'code.option' < ./options.html | \
	sed -E 's/<code.*>(.*)<\/code>/\1/' | \
	go run .
