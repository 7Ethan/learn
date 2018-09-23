const file = require('fs');

file.readFile('./readfile.js',function(err,data){
    if (err){
        return console.error(err)
    }
    console.log(data.toString())
});

console.log('over');