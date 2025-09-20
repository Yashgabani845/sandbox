#!/bin/sh

# $1 → source code filename (e.g., main.cpp)
SRC="$1"
OUT="out"
INPUT="${2:-/dev/null}"

# Compile
g++ -O2 -std=c++17 "$SRC" -o "$OUT"

# Run and save output
./"$OUT" < "$INPUT" > output.txt
