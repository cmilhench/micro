const fs = require('fs');
const path = require('path');

const name = 'Fibonacci Service'
let revision = 'unknown';
try {
    revision = fs.readFileSync(path.join(__dirname, '..', '.version'), 'utf8').trim();
} catch (error) {
    console.warn('Warning: .version file not found or unreadable');
}

const ServiceName = () => `${name} ${revision}`;


const Fibonacci = (n) => {
    if (n <= 1) return n;
    return Fibonacci(n-1) + Fibonacci(n-2);
};

const Sequence = () => {
    let x = 0, y = 1;
    return () => {
        const temp = x + y;
        x = y;
        y = temp;
        return x;
    };
};

module.exports = {
    ServiceName,
    Fibonacci,
    Sequence
};
