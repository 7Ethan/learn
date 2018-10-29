var fs = require('fs');

var readStream = fs.createReadStream('ReadWriteStream.js');

var writeStream = fs.createWriteStream('output.js');

readStream.pipe(writeStream)

console.log("程序执行完毕");

// Node.js 提供了 exports 和 require 两个对象，
// 其中 exports 是模块公开的接口，
// require 用于从外部获取一个模块的接口，即所获取模块的 exports 对象