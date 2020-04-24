NodeJS

* Single thread
* Event Loop (Events proccessed one at a time).

### Blocking Code

Read file from filesystem, set equal to contents
Print contents
Do Something else

        var contents = fs.readFileSync('/etc/hosts');
        console.log(contents);
        console.log("Do Something else");

### Non-Blocking code

Read file from filesystem
   Whenever it is complete, Print contents (callback)
Do Something else


        fs.readFile('/etc/hosts)',function(err,contents){
            console.log(contents);
        });
        console.log("Do Something else");

To build following kind of appllications, node suits the best

* Websocket server (chat server)
* Fast file upload client
* Ad Server
* Any Real-Time Data Apps

Simple HTTP Server

        # hello.js

        var http = require('http');
        http.createServer(function(request, response){
            response.writeHead(200);
            response.write("Hello, this is http Server");
            response.end();
        }).listen(8080);

        console.log('listening on port 8080...');

        # running the application
        $ node hello.js

        # Two callbacks

        var http = require('http');
        http.createServer(function(request, response){
            response.writeHead(200);
            response.write("Hello, this is http Server");
            setTimeOut(function(){
                response.write("timeout called");
                response.end("Call end");
            },5000);
        }).listen(8080);

Note: setTimeOut doesn't block the system.

### Events

* net.Server (EventEmitter)
* fs.readStream (EventEmitter)

        var EventEmitter = require('events').EventEmitter;

        var logger = new EventEmitter(); // error, warn, info

        logger.on('error',function(message){
            console.log('ERR: ' + message);
        });

        logger.emit('error','No users found');

        logger.emit('error','Some error');

        var server = http.createServer();
        server.on('request',function(request,response){

        });

        server.on('close',function(request,response){
            
        });


        var events = require('events');
        var EventEmitter = events.EventEmitter;

        var chat = new EventEmitter();
        var users = [], chatlog = [];

        chat.on('message', function(message) {
        chatlog.push(message);
        });

        chat.on('join', function(nickname) {
        users.push(nickname);
        });

        // Emit events here
        chat.emit('join','Finch');
        chat.emit('message','Hello Machine!');

We can add more than  1 listener to same event

        server.on('request', function(request, response) {
            response.writeHead(200);
            response.write("Hello, this is dog");
            response.end();
        });

        server.on('request', function(request, response) {
            console.log("New request coming in...");
        });


        // does not accepts paramaters 
        server.on('close', function() {
            console.log("Closing down the server...");
        });

### Streams

Streams are like channels which they can be readable,writable, or both

in request event of http.CreateServer

request is a readable stream
response is a writable stream

* Readable Stream(EventEmitter) -- readable, end(Events)


        var http = require('http');
        http.createServer(function(request, response){
            response.writeHead(200);
            request.on('readable',function(){
                var chunk = null;
                while( null !== (chunk = request.read())){
                    console.log(chunk.toString());//chunk here is binary
                    response.write(chunk);;// no need of toString  write does it
                }
            });
            request.on('end',function(){
                response.end();
            });
        }).listen(8080);

        request.pipe(response); // combines both readble and end events of above code

        $ curl -d 'hello' http://localhost:8080
        hello

        var fs = require('fs');
        var file = fs.createReadStream("readme.md");
        var newFile = fs.createWriteStream("readme_copy.md");
        file.pipe(newFile);

        http.createServer(function(request,response){
            var newFile = fs.createWriteStream("readme_copy.md");
            request.pipe(newFile);

            request.on('end',function(){
                response.end('uploaded!');
            });
        }).listen(8080);

        $ curl --upload-file  readme.md http://localhost:8080
        uploaded!


        http.createServer(function(request,response){
            var newFile = fs.createWriteStream("readme_copy.md");
            var fileBytes = request.headers['content-length'];
            var uploadedBytes = 0;
            request.on('readable',function(){
                var chunk = null;
                while( null !== (chunk = request.read())){
                    uploadeBytes += chunk.length;
                    var progress = (uploadedBytes / fileBytes ) * 100;
                    response.write("progress " + parseInt(progress,10) + "%\n");
                }
            });

            request.pipe(newFile);

            request.on('end',function(){
                response.end('uploaded!');
            });
        }).listen(8080);



        var fs = require('fs');

        var file = fs.createReadStream('origin.txt');
        var destFile = fs.createWriteStream('destination.txt');

        // without second parametes it throws error
        // as pipe closes the write stream automatically
        file.pipe(destFile, { end: false });

        file.on('end', function () {
        destFile.end('Finished!');
        });

        var fs = require('fs');
        var http = require('http');

        http.createServer(function(request, response) {
        response.writeHead(200, {'Content-Type': 'text/html'});

        var file = fs.createReadStream('index.html');
            file.pipe(response);
        }).listen(8080);



