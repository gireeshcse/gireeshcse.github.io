/**
 * This file generates html from markdown
 */

var marked = require('marked');
var fs = require('fs');

var arguments = process.argv;
// arguments[0] => node
// arguments[1] => generateHTML.js
// arguments[2] => markdown file
// arguments[3] => optional filename to be generated

if(process.argv.length != 3){
    console.log("Invalid Execution");
    process.exit(1);//Failure
}

try {
    if (fs.existsSync(arguments[2])) {
      //file exists
        var mdFile = fs.readFileSync(arguments[2], 'utf-8');
        var mdFileOutput = marked(mdFile);
        fs.writeFileSync(arguments[2]+'.html', mdFileOutput);
        console.log(arguments[2]+'.html is written');
    }else{
        console.log(arguments[2]+ " File not found exists");
        process.exit(1);//Failure
    }
} 
catch(err) {
    console.log("markdown file not found");
    console.error(err);
}

process.exit();//Success
