const http = require('http');

const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer((req, res) => {
    let body = '<h1>Hello World</h1>';

    res.writeHead(200, {
        'Content-Length': body.length,
        'Content-Type': 'text/html',
        'Pesan-Header': 'Pengenalan Node.js'
    });

    res.write(body)
    res.end();
});

server.listen(port, hostname, () => {
    console.log(`Server running at http://${hostname}:${port}`)
})