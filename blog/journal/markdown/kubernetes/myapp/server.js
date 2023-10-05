const http = require("http");
const os = require("os");

const listenPort = 8080;

var handler = function(request,response){
    let clientIP = request.connection.remoteAddress;
    response.writeHead(200);
    response.write("Hostname "+os.hostname())
    response.write("\nYour IP "+clientIP)
    response.write("\nRequest url "+request.url)
    response.end("\n");
}

console.log("Nodejs server is starting...\n");
console.log("Local Hostname is "+os.hostname());
console.log("\nListening on port "+listenPort)
var server = http.createServer(handler);
server.listen(listenPort);