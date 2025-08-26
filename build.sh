#!/usr/bin/env nix-shell
#!nix-shell -i bash
#!nix-shell -p htmlq -p gnused

htmlq 'code.option' < ./options.html > ./options.txt
sed -E 's/<code.*>(.*)<\/code>/\1/' options.txt > options.txt.tmp
cp options.txt.tmp options.txt
rm options.txt.tmp

