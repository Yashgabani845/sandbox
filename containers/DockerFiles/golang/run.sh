#!/bin/bash
# Usage: ./run.sh <code.go> <input.txt> <output.txt>

if [ $# -ne 3 ]; then
  echo "Usage: $0 <code.go> <input.txt> <output.txt>"
  exit 1
fi

CODE=$1
INPUT=$2
OUTPUT=$3

# Make sure output directory exists
OUTDIR=$(dirname "$OUTPUT")
mkdir -p "$OUTDIR"

docker run --rm \
  -v $(pwd)/$CODE:/runner/main.go \
  -v $(pwd)/$INPUT:/runner/input.txt \
  -v $(pwd)/$OUTDIR:/runner \
  go-executor

echo "✅ Output written to $OUTPUT"
