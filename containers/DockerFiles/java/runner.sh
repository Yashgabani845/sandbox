#!/bin/sh
SRC="$1"
INPUT="${2:-/dev/null}"
CLASSNAME=$(basename "$SRC" .java)

# Compile
javac "$SRC" 2> compile.err
if [ $? -ne 0 ]; then
    echo "Compilation Error:"
    cat compile.err
    exit 1
fi

# Run
java "$CLASSNAME" < "$INPUT"
