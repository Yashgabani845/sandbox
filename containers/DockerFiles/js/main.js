// main.js
const fs = require('fs');

// Read from stdin
let input = '';
process.stdin.on('data', chunk => input += chunk);
process.stdin.on('end', () => {
    let lines = input.trim().split('\n');
    let n = parseInt(lines[0]);
    console.log(n * 2); // simple doubling
});
