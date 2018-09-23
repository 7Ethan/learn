var fs = require('fs');
var data = '';

var readerStream = fs.createReadStream('stream.js');

readerStream.setEncoding('utf-8');

readerStream.on('data',function(chunk){
    data += chunk;
})

readerStream.on('error',function(err){
    console.error(err.stack)
})

readerStream.on('end',function(){
    console.log("it is over : data is " + data);
})

console.log("程序执行完毕");