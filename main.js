let httpServer = require('http-server')
let tailwindcss = require('tailwindcss')


let hs = httpServer.createServer({
    root: '.'
})

hs.listen(8000)