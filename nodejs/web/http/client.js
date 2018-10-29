var http = require('http');

var option = {
    host:"127.0.0.1",
    port:3000,
    path:'/index.html'
}

var callback = function(resp){
    var body = '';
    resp.on('data',function(data){
        body += data;
    })

    resp.on('end',function(){
        console.log(body);
    })
}

http.request(option,callback).end();

