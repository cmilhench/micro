const express = require('express');
const app = express();

const fib = require('./fib');

const safeParseInt = (value, defaultValue = null) => {
    const parsed = parseInt(value);
    return isNaN(parsed) ? defaultValue : parsed;
};

const createRoutes = () => {
    const router = express.Router();

    router.get('/v', (req, res) => {
        res.send(fib.ServiceName() + '\n');
    });

    router.get('/fib/n/:num', (req, res) => {
        const num = safeParseInt(req.params.num);
        if (num === null) {
            res.status(404).end();
            return;
        }
        res.setHeader('Content-Type', 'text/plain; charset=utf-8');
        res.send(fib.Fibonacci(num).toString() + '\n');
    });

    router.get('/fib/s/:num?', (req, res) => {
        const num = safeParseInt(req.params.num, 10);
        const next = fib.Sequence();
        
        res.setHeader('Content-Type', 'text/plain; charset=utf-8');
        let response = '';
        for (let i = 0; i < num; i++) {
            response += next().toString() + '\n';
        }
        res.send(response);
    });

    router.use('*', (req, res) => {
        res.status(404).send('404 Not Found\n');
    });
    
    return router;
};

const main = () => {
    const port = process.env.PORT || '8080';
    
    const app = express();
    app.use('/', createRoutes());

    app.listen(port, () => {
        console.log(`Starting ${fib.ServiceName()} :${port}`);
    }).on('error', (err) => {
        console.error(err);
    });
};

main();

module.exports = { createRoutes };
