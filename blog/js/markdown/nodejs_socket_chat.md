### Web Sockets

 npm install --save socket.io


    // server is http server
    var io = require('socket.io')(server);
    io.on('connection',function(client){
        console.log('Client connected...');
        client.emit('messages',{hello: 'Hello World'});

        client.on('messages',function(message){
            console.log(message);
        });
        /// broadcast to all the users
        client.on('sendToAll',function(data){
            client.broadcast.emit("messages",data);
        });
        client.on('join',function(name){
            // now accessible for all listeners
            client.nickname = name;
        });
    });

In client

    <script src="/socket.io/socket.io.js"></script>
    <script>
        var socket = io.connect('http://localhost:8080');
        socket.on('message',function(data){
            console.log(data.hello);
        });

         socket.on('messages',function(data){
           insertMessage(data);
        });

        $('#chat_form').submit(function(e){
            var message = $('#chat_input').val();
            socket.emit('messages',message);
        });

        socket.emit('join',nickname); // notify the server
    </script>


app.js

        let app = express();

        var server = http.createServer(app);
        server.listen(8080);

        var io = require('socket.io')(server);

        io.sockets.on('connection',function(client){
            console.log('Client connected...');
        });

client

        <script src='/socket.io/socket.io.js'></script>

        <script>
        // use the socket.io server to connect to localhost:8080 here
        var socket = io.connect('http://localhost:8080');
        </script>


client 

        <script src="/socket.io/socket.io.js"></script>

        <script>
        var insertQuestion = function (question) {
        var newQuestion = document.createElement('li');
        newQuestion.innerHTML = question;

        var questions = document.getElementsByTagName('ul')[0];
        return questions.appendChild(newQuestion);
        }

        let server = io.connect('http://localhost:8080');

        server.on('question',function(question){
            insertQuestion(question);
        });
        </script>

Server

        let express = require('express');
        let app = express();
        let server = require('http').createServer(app);
        let io = require('socket.io')(server);

        io.on('connection', function(client) {
        console.log("Client connected...");
        
        client.on('question',function(question){
            client.broadcast.emit('question',question);
        });
        
        
        });

        server.listen(8080);

Server another variation

        let express = require('express');
        let app = express();
        let server = require('http').createServer(app);
        let io = require('socket.io')(server);

        io.on('connection', function (client) {
        console.log("Client connected...");

        client.on('question', function (question) {
            //only one question can be asked
            if(client.question_asked != true)
            {
            client.question_asked = true;
            client.broadcast.emit('question',question);
            }
            
        });
        });

        server.listen(8080);


Version 2

        let express = require('express');
        let app = express();
        let server = require('http').createServer(app);
        let io = require('socket.io')(server);

        io.sockets.on('connection', function(client) {
        console.log("Client connected...");

        // listen for answers here
        client.on('answer',function(question,answer){
            client.broadcast.emit('answer',question,answer);
        });
        

        client.on('question', function(question) {
            if(!client.question_asked) {
            client.question_asked = true;
            client.broadcast.emit('question', question);
            }
        });
        });

        server.listen(8080);

Client

        <script src="/socket.io/socket.io.js"></script>

        <script>
        let server = io.connect('http://localhost:8080');

        server.on('question', function (question) {
            insertQuestion(question);
        });
        
        server.on('answer',function(question,answer){
            answerQuestion(question,answer);
        });

        //Don't worry about these methods, just assume 
        //they insert the correct html into the DOM
        // let insertQuestion = function(question) {
        // }

        // let answerQuestion = function(question, answer) {
        // }
        </script>

### Persisting Data

        var messages = [];
        var storeMessage = function(name, data){
            messages.push({name:name, data:data});
            if(messages.length > 10)
            {
                messages.shift();// remove first message
            }
        };

        io.sockets.on('connection',function(client){
            client.on("messages",function(message){
                client.get("nickname",function(error,name){
                    client.broadcast.emit("messages",name+' '+message);
                    storeMessage(name,message);
                });
            });

            client.on('join',function(name){
                client.nickname = name;
                messages.forEach(function(message){
                    client.emit("messages", message);
                });
            });
        });

#### redis

Redis is a key-value storage(in-memory data structure store)

Used as

* Database
* Cache and Message broker

Supports datastructures such as

* Strings
* Hashes
* Lists
* Sets
* Sorted sets with range queries
* Bitmaps
* Hyperlogs
* Geospatial indexes with radius queries
* Streams

Some Commands

        SET server:name "Backend Server"
        GET server:name  => "Backend Server"

        EXISTS server:name => 1
        EXISTS server:bla => 0

        SET connections 10
        INCR connections => 11
        INCR connections => 12
        DEL connections
        INCR connections => 1

        INCRBY connections 10 => 11

        DECR connections => 10
        DECRBY connection 5 => 5

Note: INCR is an atomic operation.

**All the Redis operations implemented by single commands are atomic, including the ones operating on more complex data structures. So when you use a Redis command that modifies some value, you don't have to think about concurrent access.**

        SET resource:lock "Redis Demo"
        EXPIRE resource:lock 120

        TTL resource:lock => 113
        TTL resource:lock => -2

The -2 for the TTL of the key means that the key does not exist (anymore). A -1 for the TTL of the key means that it will never expire.

         # the below command is a atomic operation
         SET resource:lock "Redis Demo 3" EX 5
         TTL resource:lock => 5
         # to remove TTL
         PERSIST resource:lock

#### List Commands

* RPUSH

Puts new element at the end of the list.

        RPUSH friends "Ram"
        RPUSH friends 1 2 3
        RPUSH friends "Vedavati" "Ravana"  "Hanuman"

* LPUSH

Puts new element at the star of the list.

        LPUSH friends "Shiva"
        LPUSH "Sita" "Aryavartha"

* LLEN

        LLEN friends => 2

* LRANGE

Gives subset of the list.

        RPUSH friends "Sam"
        RPUSH friends "Alice"
        RPUSH friends "Bob"
        LRANGE friends 0 -1 => 1) "Sam", 2) "Alice", 3) "Bob"
        LRANGE friends 0 1 => 1) "Sam", 2) "Alice"
        LRANGE friends 1 2 => 1) "Alice", 2) "Bob"
        LRANGE friends 0 -3 => 1) "Sam"



* LPOP

removes the first element from the list and returns it.

        LPOP friends => "Sam"

* RPOP

 removes the last element from the list and returns it.

        RPOP friends => "Bob"

#### Set commands

* SADD
* SREM
* SISMEMBER
* SMEMBERS
* SUNION
* SPOP

     SADD letters a b c d e f => 6
     SPOP letters 1) "c"
     SPOP letters 1) "e"

* SRANDMEMBER

return random elements without removing such elements from the set

    SRANDMEMBER letters 6 => 1) "f" 2) "a" 3) "d" 4) "b"
    SRANDMEMBER letters => 1) "f"

#### Sorted Sets

A sorted set is similar to a regular set, but now each value has an associated score. This score is used to sort the elements in the set.

* ZADD
* ZSCORE
* ZREM
* ZRANK
* ZRANGE
* ZRANGEBYSCORE
* ZREMRANGEBYSCORE
* ZREVRANGE

    ZADD hackers 1940 "Alan Kay"
    ZADD hackers 1906 "Grace Hopper"
    ZADD hackers 1953 "Richard Stallman"
    ZADD hackers 1965 "Yukihiro Matsumoto"
    ZADD hackers 1916 "Claude Shannon"
    ZADD hackers 1969 "Linus Torvalds"
    ZADD hackers 1957 "Sophie Wilson"
    ZADD hackers 1912 "Alan Turing"

    ZRANGE hackers 2 4 => 1) "Claude Shannon", 2) "Alan Kay", 3) "Richard Stallman"

    ZRANK hackers "Linus Torvalds" => 7
    ZREM hackers "Alan Turing" => 1
    ZSCORE hackers "Linus Torvalds" => 1969.0
    
#### Hashes

Hashes are maps between string fields and string values

    HSET user:1000 name "John Smith"
    HSET user:1000 email "john.smith@example.com"
    HSET user:1000 password "s3cret"

    HGETALL user:1000

    HMSET user:1001 name "Mary Jones" password "hidden" email "mjones@example.com"

    HGET user:1001 name

    HSET user:1000 visits 10
    HINCRBY user:1000 visits 1 => 11
    HINCRBY user:1000 visits 10 => 21
    HDEL user:1000 visits
    HINCRBY user:1000 visits 1 => 1


    npm install redis

    var redis = require('redis');
    var client = redis.createClient();

    client.set('name','finch');

    client.get('name',function(err,data){
        console.log(data);
    });

    let question1 = "Where is the dog?";
    let question2 = "Where is the cat?";

    client.lpush("questions",question1, function(err,response){
        console.log(response);
    });

    client.lpush("questions",question2, function(err,response){
        console.log(response);
    });

    client.lrange("questions",0,-1,function(err,questions){
        console.log(questions);
    });


Example 

        let express = require('express');
        let app = express();
        let server = require('http').createServer(app);
        let io = require('socket.io').listen(server);

        let redis = require('redis');
        let redisClient = redis.createClient();

        io.sockets.on('connection', function (client) {

        // enter your code here
        redisClient.lrange("questions",0,-1,function(err,questions){
            questions.forEach(function(question){
            client.emit('question',question);
            });
        });

        client.on('answer', function (question, answer) {
            client.broadcast.emit('answer', question, answer);
        });

        client.on('question', function (question) {
            if (!client.question_asked) {
            client.question_asked = true;
            client.broadcast.emit('question', question);
            redisClient.lpush("questions", question);
            }
        });
        });


        client.on('question', function (question) {
            if (!client.question_asked) {
            client.question_asked = true;
            client.broadcast.emit('question', question);
            redisClient.lpush("questions", question,function(err,response){
                redisClient.ltrim("questions",0,19); // keeps newest 20 questions
            });
            }
        });
 

## RabbitMQ

RabbitMQ is open source message broker software (sometimes called message-oriented middleware) that implements the Advanced Message Queuing Protocol (AMQP). 

    docker pull rabbitmq:3
    docker run -d --hostname my-rabbit --name some-rabbit rabbitmq:3

