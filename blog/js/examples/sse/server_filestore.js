const express = require('express');
const EventEmitter = require('events');

// For File storage
const fs = require('fs');
const readline = require('readline');
const filename = `${__dirname}/messages.txt`;
const writeStream = fs.createWriteStream(filename,{flags:'a+'});

const chatEmitter = new EventEmitter();

const app = express();

const port = process.env.port | 4000;

app.use(express.static('public'));

app.get('/sse',function(req,res){
    res.writeHead(200,{
        'Content-Type': 'text/event-stream',
        'Connection': 'keep-alive'
    })

    // Reading from from and sending to the user
    var rl = readline.createInterface({
        input: fs.createReadStream(filename),
        crlfDelay: Infinity
    });
    
    rl.on('line',(line)=>{
        res.write(`data: ${line}\n\n`);
        // writing to the file
        writeStream.write(`${message} \r\n`);
    });

    rl.on('close',()=>{
        // Closed after reading on the lines doenot watch for the changes in file
        console.log('read line closed');
    });

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