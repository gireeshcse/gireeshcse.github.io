var express = require('express');

var app = express();

app.use('/',express.static(__dirname + '/../'));
console.log(__dirname + '/../');

app.get('/',(req,res)=> res.send("Hello World!"));

var server = app.listen(9000,()=> console.log(`app lisenting on port ${9000}`));