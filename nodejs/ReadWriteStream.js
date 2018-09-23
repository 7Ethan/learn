var fs = require('fs');

var readStream = fs.createReadStream('ReadWriteStream.js');

var writeStream = fs.createWriteStream('output.js');

readStream.pipe(writeStream)

console.log("程序执行完毕");