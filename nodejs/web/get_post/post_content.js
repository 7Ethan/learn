var http = require('http');
var querystring = require('querystring');

var postHTML = 
  '<html><head><meta charset="utf-8"><title> Node.js 实例</title></head>' +
  '<body>' +
  '<form method="post">' +
  '网站名： <input name="name"><br>' +
  '网站 URL： <input name="url"><br>' +
  '<input type="submit">' +
  '</form>' +
  '</body></html>';

http.createServer(function(req,resp){
    var body = '';
    req.on('data',function(chunk){
        body += chunk;
    });

    req.on('end',function(){
        // 解析参数
        body = querystring.parse(body);
        // 设置响应头部信息及编码
        resp.writeHead(200,{"Content-Type":"text/html;charset=utf8"});
        if (body.name && body.url) {
            resp.write("Name : " + body.name);
            resp.write("</br>");
            resp.write("URL : " + body.url);
        } else {
            resp.write(postHTML)
        }
        resp.end();
    });
}).listen(3000);