var http = require("http")
http.createServer(function(request,response){
    // Closeure function
    response.writeHead(200,{"Content-type":"text/plain"});
    response.write("Hello Node.js");
    response.end();
}).listen(8088);