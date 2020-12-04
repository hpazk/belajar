import { readdir } from 'fs';

readdir('./', 'utf-8', (err, content) => {
    if (err)
        return err;
    console.log(content)
});