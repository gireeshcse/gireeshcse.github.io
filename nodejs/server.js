var http = require('http');
var port = 9000;

http.createServer(function(req,res){
    res.write("Hello World!");
    console.log(req.url);
    res.end();
}).listen(port);