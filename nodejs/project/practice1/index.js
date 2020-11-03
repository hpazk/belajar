const express = require('express');

const port = process.env.PORT || 8080;

const app = express();

app.get('/', (req, res) => res.send('Page Index - Edited'));

app.listen(port, () => {
    console.log(`Server running in localhost:${port}`)
})