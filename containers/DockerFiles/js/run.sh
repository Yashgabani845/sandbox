#!/bin/sh

# $1 = JS source file (main.js)
# $2 = optional input file
SRC="$1"
INPUT="${2:-/dev/null}"

# Run JS code and redirect output
node "$SRC" < "$INPUT" > output.txt
