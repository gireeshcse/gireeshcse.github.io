const express = require('express');
const EventEmitter = require('events');


const chatEmitter = new EventEmitter();

const app = express();

const port = process.env.port | 4000;

app.use(express.static('public'));

app.get('/sse',function(req,res){
    res.writeHead(200,{
        'Content-Type': 'text/event-stream',
        'Connection': 'keep-alive'
    })

    const onMessage = (message) => {
        res.write(`data: ${message}\n\n`);
    };

    chatEmitter.on('message',onMessage);

    // This prevents writing messages to a closed connection
    res.on('close',()=>{
        chatEmitter.off('message',onMessage)
    })

})

app.get('/chat',function(req,res){
    const { message } = req.query;
    chatEmitter.emit('message',message)
    res.end()
})


app.listen(port,()=>{
    console.log(`Server running at http://localhost:${port}`)
})