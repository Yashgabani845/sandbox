#!/bin/bash

# Usage: ./run.sh <file.c|file.cpp> [input_file]

SRC="$1"
INPUT="${2:-/dev/null}"  # stdin defaults to empty

# Determine language and compiler
EXT="${SRC##*.}"
OUT="a.out"

case "$EXT" in
    c)
        COMPILER="gcc"
        ;;
    cpp)
        COMPILER="g++"
        ;;
    *)
        echo "Unsupported file type: $EXT" >&2
        exit 1
        ;;
esac

# Compile
$COMPILER -O2 -Wall -Wextra -std=c++17 "$SRC" -o "$OUT" 2> compile_errors.txt
COMPILE_EXIT=$?

if [ $COMPILE_EXIT -ne 0 ]; then
    cat compile_errors.txt
    exit $COMPILE_EXIT
fi

# Execute with stdin
./$OUT < "$INPUT"
EXIT_CODE=$?

# Cleanup
rm -f "$OUT" compile_errors.txt

# Return exit code
exit $EXIT_CODE
