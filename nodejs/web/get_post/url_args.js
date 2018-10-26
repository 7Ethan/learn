var http = require('http');
var url = require('url');
var util = require('util');

http.createServer(function(req,resp){
    resp.writeHead(200,{"Content-Type":"text/plain"});
    var params = url.parse(req.url,true).query;
    resp.write("Name is :" + params.name);
    resp.write("\n");
    resp.write("Email is :" + params.email);
    resp.end();
}).listen(3000);