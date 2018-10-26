var http = require('http');
var url = require('url');


// similar ti golang's middware
function start (route) {
    function onRequest(request,response){
        var pathname = url.parse(request.url).pathname;
        console.log('Request is querying' + pathname );

        route(pathname);
        
        response.writeHead(200, {"Content-Type": "text/plain"});
        response.write("Hello World");
        response.end()
    }
    
    http.createServer(onRequest).listen(8080);
    console.log('Server has been started !')
}

exports.start = start;