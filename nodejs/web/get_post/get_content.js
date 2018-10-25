var http = require('http');
var url = require('url');
var util = require('util');

http.createServer(function(req,resp){
    resp.writeHead(200,{"Content-Type":"text/plain"});
    resp.end(util.inspect(url.parse(req.url,true)));
}).listen(3000);