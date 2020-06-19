# Node.js

* All I/O tasks are non-blocking and asynchronous

#### Applications

* HTTP APIs
* Distributed systems
* command-line tools, and
* cross-platform desktop applications.

#### Not best choice for

* CPU intensive applications (Low level languages like C, Go, Rust etc are best for these).

#### Popular frameworks

* express
* hapi
* restify

#### Examples

```
require('http')
    .createServer((req, res) => res.end('hello world!'))
    .listen(8080)

```

```
res.setHeader('Content-Type', 'application/json');
res.end(JSON.stringify({ 'name': 'ram', brothers: ['Lakshmana','Bharata','Shatrughna'] }));

res.setHeader('Content-Type', 'text/plain')
res.end('India')

res.writeHead(404, { 'Content-Type': 'text/plain' })
res.end('Not Found')

req.url         // www.google.com?data=20&id=10
req.url.split('?') // ["www.google.com", "data=20&id=10"]
       .slice(1)  // ["data=20&id=10"]
       .join('') // "data=20&id=10"

const querystring = require('querystring'); // core library

const obj = querystring.parse('foo=bar&abc=xyz&abc=123'); 
/*
{
  foo: 'bar',
  abc: ['xyz', '123']
}
*/

req.url.match(/^\/public/) # matches /public
```

Creating a Stream Object and loads data from filesystem and sends it to the client via response object.

```
const fs = require('fs');
const filename = `${__dirname}/downloads${req.url.split('/public')[1]}`;
fs.createReadStream(filename)
    .on('error',() => {
        res.writeHead(404, { 'Content-Type': 'text/plain' })
        res.end('File not found')
    })
    .pipe(res)

```

#### Express

* Fast, unopinionated, minimalist web framework for Node.js


```
const express = require('express');
const app = express()
const port = 4000

app.get('/',(req,res) => res.send('hello world!'))

app.listen(port, () => console.log(`Server running at http://localhost:${port}`))

app.METHOD(PATH, HANDLER) // put, get, delete, post

app.post('/', function (req, res) {
  res.send('Got a POST request')
})

res.json({ 'name': 'ram', brothers: ['Lakshmana','Bharata','Shatrughna'] });

//http://localhost:4000/?name=ram&brothers=bharata&brothers=shatrughna
req.query //{"name":"ram","brothers":["bharata","shatrughna"]}

```

Serving static files

```
app.use(express.static('public'))
app.use(express.static('files'))
app.use('/static', express.static('public'))
app.use('/static', express.static(path.join(__dirname, 'public')))
app.use('/download',express.static(__dirname));

/**
http://localhost:4000/others/sample
/others/sample
*/
app.get('/others/*',function(req,res){
    res.send(req.url);
})
```

#### SSE - Server-Sent Events

SSE is available in most browsers through the **EventSource API**, and is a very simple way for the server to “push” messages to the browser.

**EventSource** instance opens a persistent connection to an HTTP server, which sends events in **text/event-stream** format. The connection remains open until closed by calling **EventSource.close()**.

Unlike WebSockets, server-sent events are unidirectional; that is, data messages are delivered in one direction, from the server to the client (such as a user's web browser). That makes them an excellent choice when there's no need to send data from the client to the server in message form.

EventSource is a useful approach for handling things like social media status updates, news feeds, or delivering data into a client-side storage mechanism like IndexedDB or web storage.

```
var evtSource = new EventSource('sse.php');
var eventList = document.querySelector('ul');

evtSource.onmessage = function(e) {
  var newElement = document.createElement("li");

  newElement.textContent = "message: " + e.data;
  eventList.appendChild(newElement);
}

const sse = new EventSource('/api/v1/sse');

  /* This will listen only for events 
   * similar to the following:
   * 
   * event: notice
   * data: useful data
   * id: someid
   *
   */
  sse.addEventListener("notice", function(e) { 
    console.log(e.data)
  })

  /* Similarly, this will listen for events 
   * with the field `event: update`
   */
  sse.addEventListener("update", function(e) {
    console.log(e.data)
  })

  /* The event "message" is a special case, as it
   * will capture events without an event field
   * as well as events that have the specific type
   * `event: message` It will not trigger on any
   * other event type.
   */
  sse.addEventListener("message", function(e) {
    console.log(e.data)
  })

```

(Events Specification)[https://html.spec.whatwg.org/multipage/server-sent-events.html#server-sent-events]

##### EventEmitter

```
const EventEmitter = require('events');

const myEmitter = new EventEmitter();

// First listener
myEmitter.on('event', function firstListener() {
  console.log('Helloooo! first listener');
});
// Second listener
myEmitter.on('event', function secondListener(arg1, arg2) {
  console.log(`event with parameters ${arg1}, ${arg2} in second listener`);
});
// Third listener
myEmitter.on('event', function thirdListener(...args) {
  const parameters = args.join(', ');
  console.log(`event with parameters ${parameters} in third listener`);
});
myEmitter.emit('event', 1, 2, 3, 4, 5);
// Helloooo! first listener
// event with parameters 1, 2 in second listener
// event with parameters 1, 2, 3, 4, 5 in third listener


```

**Example**

```
// public/chat.html
<!DOCTYPE html>
<html lang="en">
<title>Chat App</title>
<body>
<div id="messages">
    <h4>Chat Messages</h4>
</div>

<form id="form">
    <input id="input" type="text" placeholder="Your message...">
    <input type="submit" value="Send Message" />
</form>

<script type="text/javascript">
    var eventSource = new window.EventSource('/sse');

    eventSource.onmessage = function(event){
        document.getElementById('messages').innerHTML += `<p>${event.data}</p>`;
    }

    document.getElementById('form').addEventListener('submit',function(evt){
        evt.preventDefault();
        window.fetch(`/chat?message=${document.getElementById('input').value}`);
        document.getElementById('input').value = '';
    });

</script>
</body>
</html>

// part of server.js

const EventEmitter = require('events');
const chatEmitter = new EventEmitter();
app.use(express.static('public'));

app.get('/sse',function(req,res){
    res.writeHead(200,{
        'Content-Type': 'text/event-stream',
        'Connection': 'keep-alive'
    })

    const onMessage = (message) => {
        res.write(`data: ${message}\n\n`)
    };

    chatEmitter.on('message',onMessage);

    // This prevents writing messages to a closed connection
    res.on('close',function(){
        chatEmitter.off('message',onMes
        sage)
    })
})

app.get('/chat',function(req,res){
    const { message } = req.query;
    chatEmitter.emit('message',message)
    res.end()
})

```

To read and write messages to the file

```
const fs = require('fs');
const readline = require('readline');

const filename = `${__dirname}/messages.txt`;
const writeStream = fs.createWriteStream(filename,{flags:'a+'});

app.get('/sse',function(req,res){
    res.writeHead(200,{
        'Content-Type': 'text/event-stream',
        'Connection': 'keep-alive'
    })

    var rl = readline.createInterface({
        input: fs.createReadStream(filename),
        crlfDelay: Infinity
    });
    
    rl.on('line',(line)=>{
        res.write(`data: ${line}\n\n`);
    });

    rl.on('close',(line)=>{
        // Closed after reading on the lines doenot watch for the changes in file
        console.log('read line closed');
    });

    
    const onMessage = (message) => {
        res.write(`data: ${message}\n\n`);
        writeStream.write(`${message} \r\n`);
    };

    chatEmitter.on('message',onMessage);

    // This prevents writing messages to a closed connection
    res.on('close',()=>{
        chatEmitter.off('message',onMessage)
    })

})
```


#### process

* Globally available object
* **process.argv** for command line arguments
* **process.env** environmental variables of the system as well as any additions by application.
* **process.exit()** exits the node process

#### Blocking Code

```
let count = 0
// below code will not execute
setInterval(() => console.log(`${++count} count`), 1000)

setTimeoutSync(5500)
console.log('hello from the past!')

process.exit()

function setTimeoutSync (ms) {
    const t0 = Date.now()
    while (Date.now() - t0 < ms) {}
}

// Output
// hello from the past!
```

#### Files

* fs.readFileSync()
* fs.readdir() - Follows c naming covention since it uses a system call and it  will only return file names


Listen to file changes

```
fs.watchFile(filename,(curr,prev)=>{
    console.log('file changed');
    // fire event
});

```
Asynchronously reads the entire contents of a file.

```
fs.readFile('/etc/passwd', (err, data) => {
  if (err) throw err;
  console.log(data);
  console.log(data.length);
  console.log(fileData.toString())
});
```

#### Callbacks - Non Blocking Code

```
function mapAsync(arr, fn, onFinish)
{
    let prevError
    let nRemaining = arr.length
    const results = []
    
    arr.forEach(function(item,i){
        fn(item,function(err,data){
            if (prevError) return

            if(err)
            {
                prevError = err
                return onFinish(err)
            }
            results[i] = data
            // To keep track of how many items are processed
            nRemaining--
            if(!nRemaining) onFinish(null,results)
        });
    });
}

mapAsync([10,20,30,40,50],(value,fn)=>{
    return fn(null,value/10)
},(err,results)=>{
    if (err) return console.error(err)

    console.log(results)
})

mapAsync([10,20,30,40,0,50],(value,fn)=>{
    if (value === 0)
        return fn({error:'can not contain zero'})
    return fn(null,value/2);
},(err,results)=>{
    if (err) return console.error(err)

    console.log(results)
})

mapAsync(['sample.js', 'server.js'], fs.readFile, (err, filesData) => {
    if (err) return console.error(err)
    console.log(filesData)
})

console.log('mapAsync completed')

fs.readdir('./', function (err, files) {
    if (err) return console.error(err)

    mapAsync(files, fs.readFile, (err, results) => {
        if (err) return console.error(err)

        results.forEach((data, i) => console.log(`${files[i]}: ${data.length}`))
        console.log('done!')
    })
})

const targetDirectory = process.argv[2] || './'

getFileLengths(targetDirectory, function (err, results) {
    if (err) return console.error(err)

    results.forEach(([file, length]) => console.log(`${file}: ${length}`))
    console.log('done!')
})

function getFileLengths (dir, cb) {
    fs.readdir(dir, function (err, files) {
        if (err) return cb(err)
        const filePaths = files.map(file => path.join(dir, file))
        mapAsync(filePaths, readFile, cb)
    })
}

function readFile (file, cb) {
    fs.readFile(file, function (err, fileData) {
        if (err) {
            if (err.code === 'EISDIR') return cb(null, [file, 0])
            return cb(err)
        }
        cb(null, [file, fileData.length])
    })
}

```

mapAsync() expects fn to be an asynchronous function. This means that when we execute fn() for an item in the array, it won’t immediately return a result; the result will be passed to a callback

#### Promises

* A promise is an object that represents a future action and its result.
* promise is non-blocking.


```
fs.promises.readFile('sample.js')
    .then(data => console.log(`sample.js ${data.length}`))
    .catch(err => console.error(err))
```

```
fs.promises.readdir('./')
    .catch(err => console.error(err))
    .then(files => {
        files.forEach(function (file) {
        fs.promises.readFile(file)
            .catch(err => console.error(err))
            .then(data => console.log(`${file}: ${data.length}`))
    })
    console.log('done!') // printed first
})
```

* **Promise.all()** is globally available in Node.js, and it execute an array of promises at the same time.
* We use Promise.all() to execute an array of promises all in parallel.

```
const fs = require('fs').promises
fs.readdir('./')
    .then(fileList =>
        Promise.all(
            fileList.map(file => fs.readFile(file).then(data => [file, data.length]))
        )
    )
    .then(results => {
        results.forEach(([file, length]) => console.log(`${file}: ${length}`))
        console.log('done!')
    })
    .catch(err => console.error(err))

```

```
fileList.map(file => fs.readFile(file))
```
When Promise.all() finishes, results will simply be an array of file data without filenames.

```
const fs = require('fs').promises

const targetDirectory = process.argv[2] || './'

getFileLengths(targetDirectory)
    .then(results => {
        results.forEach(([file, length]) => console.log(`${file}: ${length}`))
        console.log('done!')
    })
    .catch(err => console.error(err))

function getFileLengths (dir) {
    return fs.readdir(dir).then(fileList => {
                const readFiles = fileList.map(file => {
                    const filePath = path.join(dir, file)
                    return readFile(filePath)
                })
                return Promise.all(readFiles)
            })
}

function readFile (filePath) {
    return fs.readFile(filePath)
            .then(data => [filePath, data.length])
            .catch(err => {
                if (err.code === 'EISDIR') return [filePath, 0]
                throw err
            })
}

```

In contrast to callbacks, the success and error paths of callbacks are handled with separate functions.

#### Async & Await

These allows us to use promises as if they were synchronous.

```
const fs = require('fs').promises

printLength('10-read-file-await.js')

async function printLength (file) {
    try {
        const data = await fs.readFile(file)
        console.log(`${file}: ${data.length}`)
    } catch (err) {
        console.error(err)
    }
}

printLengths('./')

async function printLengths (dir) {
    const fileList = await fs.readdir(dir)

    const results = fileList.map(
        async file => await fs.readFile(file).then(data => [file, data.length])
    )

    results.forEach(result => console.log(`${result[0]}: ${result[1]}`))

    console.log('done!')
}
// O/P

node 11-read-dir-await-broken.js
undefined: undefined
undefined: undefined

```

Instead of printing the length of a file, we’re printing the length of a promise – which is undefined .

```
const fs = require('fs').promises

printLengths('./')

async function printLengths (dir) {
    const fileList = await fs.readdir(dir)

    const results = await Promise.all(
        fileList.map(file => fs.readFile(file).then(data => [file, data.length]))
    )
    
    results.forEach(([file, length]) => console.log(`${file}: ${length}`))
    console.log('done!')
}
```

```
const targetDirectory = process.argv[2] || './'

printLengths(targetDirectory)

async function printLengths (dir) {
    try {
        const results = await getFileLengths(dir)
        results.forEach(([file, length]) => console.log(`${file}: ${length}`))
        console.log('done!')
    } catch (err) {
        console.error(err)
    }
}

async function getFileLengths (dir) {
    const fileList = await fs.readdir(dir)
    
    const readFiles = fileList.map(async file => {
        const filePath = path.join(dir, file)
        return await readFile(filePath)
    })

    return await Promise.all(readFiles)
}

async function readFile (filePath) {
    try {
        const data = await fs.readFile(filePath)
        return [filePath, data.length]
    } catch (err) {
        if (err.code === 'EISDIR') return [filePath, 0]
        throw err
    }
}

```

Under the hood, async/await is using promises. When we declare an async function, we’re really creating a promise.

#### Event Emitters

```
const readline = require('readline')
const rl = readline.createInterface({ input: process.stdin })
rl.on('line', line => console.log(line.toUpperCase()))

```

```
const http = require('http')
const readline = require('readline')
const querystring = require('querystring')

const rl = readline.createInterface({ input: process.stdin })

rl.on('line', line =>
    http.get(
        `http://localhost:4000/chat?${querystring.stringify({ message: line })}`
    )
)

// res is an event emitter
http.get('http://localhost:4000/sse',res => {
    res.on('data', data => console.log(data.toString()))
})

```
* Custom Events

```
function createEventSource (url) {
    const source = new EventEmitter()

    http.get(url, res => {
        res.on('data', data => {
            const message = data
            .toString()
            .replace(/^data: /, '')
            .replace(/\n\n$/, '')

            source.emit('message', message)
        })
    })

    return source
}

const source = createEventSource('http://localhost:4000/sse')

source.on('message', console.log)

// Other events
const eventType = message.match(/\?$/) ? 'question' : 'statement'
source.emit(eventType, message)


source.on('question', q => console.log(`Someone asked, "${q}"`))

```

#### Streams

* Streams are a specific type of event emitter.

```
http.get(url, res => {
    res.on('data', data => {
        // here's where we use the data
    })
})
```

* **res** is a stream, and because streams are event emitters, we can use on() to listen to its events and receive updates.
* “data”, “error”, and “end” events

```
const fs = require('fs')
const https = require('https')

const fileUrl =
'https://www.fullstackreact.com/assets/images/fullstack-react-hero-book.png'

https.get(fileUrl, res => {
    const chunks = []
    res.on('data', data => chunks.push(data)).on('end', () =>
            fs.writeFile('book.png', Buffer.concat(chunks), err => {
            if (err) console.error(err)
            console.log('file saved!')
        })
    )
})
```

* **Buffer.concat()** can convert an array of Buffer objects into a single Buffer
* **fs.writeFile()** is happy to accept a Buffer as its argument for what should be written to a file.

The downside of batching is that we need to be able to hold all of the data in memory. This is not a problem for smaller files, but we can run into trouble when working with large files – especially if the file is larger than our available memory. Often, it’s more efficient to write data as we receive it.

```
const writeStream = fs.createWriteStream('time.log')
setInterval(() => writeStream.write(`The time is now: ${new Date()}\n`), 1000)

res.on('data', data => fileStream.write(data)).on('end', () => {
    fileStream.end()
    console.log('file saved!')
}

```

**pipe()**

we can connect readable streams to writable streams via pipe() .

```
https.get(fileUrl, res => {
    res
        .pipe(fs.createWriteStream('book.png'))
        .on('finish', () => console.log('file saved!'))
})
```

##### Transform Streams

A transform stream behaves as both a readable stream and a writable stream.

For example, if we wanted to efficiently transform the text in a file to upper case, we could do this:

```
const fs = require('fs')
const { Transform } = require('stream')

fs.createReadStream('notes.md')
    .pipe(shout())
    .pipe(fs.createWriteStream('loud-code.txt'))

function shout () {
    return new Transform({
        transform (chunk, encoding, callback) {
            callback(null, chunk.toString().toUpperCase())
        }
    })
}
```

Our transform stream is created by a simple function that expects three arguments. The first is the chunk of data, and we’re already familiar with this from our use of on('data') . The second is encoding, which is useful if the chunk is a string. However, because we have not changed any default behaviors with or read stream, we expect this value to be “buffer” and we can ignore it. The final argument is a callback to be called with the results of transforming the chunk. The value provided to the callback will be emitted as data to anything that is reading from the transform stream. The callback can also be used to emit an error.

**npm install csv-parser**

```
const fs = require('fs')
const csv = require('csv-parser')

fs.createReadStream('people.csv')
    .pipe(csv())
    .on('data', row => console.log(JSON.stringify(row)))

// O/P
{"name":"Liam Jones","dob":"1988-06-26"}
{"name":"Maximus Rau","dob":"1989-08-21"}
```
Add another transform stream into the mix to convert the objects to String or Buffer types as Streams are designed to work with these.

```
const YEAR_MS = 365 * 24 * 60 * 60 * 1000

fs.createReadStream('people.csv')
    .pipe(csv())
    .pipe(clean())
    .on('data', row => console.log(JSON.stringify(row)))

function clean () {
    return new Transform({
        objectMode: true, // Since data is not of String or Buffer types
        transform (row, encoding, callback) {
            const [firstName, lastName] = row.name.split(' ')
            const age = Math.floor((Date.now() - new Date(row.dob)) / YEAR_MS)
            callback(null, {firstName, lastName, age})
        }
    })
}

```