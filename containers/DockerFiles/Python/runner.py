import sys, subprocess, tempfile

# Read code from stdin
code = sys.stdin.read()

with tempfile.NamedTemporaryFile(suffix=".py", delete=False) as tmp:
    tmp.write(code.encode())
    tmp.flush()
    result = subprocess.run(
        ["python3", tmp.name],
        capture_output=True,
        text=True
    )
    print(result.stdout, end="")
    print(result.stderr, end="", file=sys.stderr)
