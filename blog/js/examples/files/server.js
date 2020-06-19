const fs = require('fs');
const readline = require('readline');

const filename = `${__dirname}/log.txt`;

const writeStream = fs.createWriteStream(filename,{flags:'a+'});
const readStream = fs.createReadStream(filename);

fs.watchFile(filename,(curr,prev)=>{
    console.log('file changed');
    // fire event
});

var rl = readline.createInterface({
    input: readStream,
    crlfDelay: Infinity
});

rl.on('line',(line)=>{
    console.log(line);
    writeStream.write(`${line} \r\n`);
});

rl.on('close',()=>{
    // Closed after reading on the lines doenot watch for the changes in file
    console.log('read line closed');
    readStream.close();
    writeStream.close();
});

// Asynchronously reads the entire contents of a file.
fs.readFile('/etc/passwd', (err, data) => {
  if (err) throw err;
  console.log(data);
});