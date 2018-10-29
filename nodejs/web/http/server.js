var http = require('http');
var fs = require('fs');
var url = require('url');

http.createServer(function(req,resp){
    var pathName = url.parse(req.url).pathname
    console.log("Request for " + pathName + " had received!");
    fs.readFile(pathName.substring(1),function(err,data){
        if (err){
            console.log("Error in reading from files:" + err)
            resp.writeHead(404,{"Content-Type":"text/html"});
            // resp.writeHead(502,{"Content-Type":"text/html"});
        } else {
            console.log("ok!");
            resp.writeHead(200,{"Content-Type":"text/html"});
            resp.write(data);
        }

        resp.end();
    })
    // resp.end();//It Can't write Here!!!!
    // events.js:183
    //   throw er; // Unhandled 'error' event
    //   ^
    //     Error: write after end
    //         at write_ (_http_outgoing.js:622:15)
}).listen(3000);

console.log("Start server on 'http://localhost:3000'")