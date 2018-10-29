var express = require('express');
var app = express();
var bodypaser = require('body-parser');


// 创建 application/x-www-form-urlencoded 编码解析
var urlencodedParser = bodypaser.urlencoded({extended:false})

app.use(express.static('public'));

app.get('/index.html',function(req,resp){
    resp.sendFile(__dirname + '/' + 'index.html');
})

app.post("/process_post",urlencodedParser,function(req,resp){
    var response = {
        "First Name : " : req.body.first_name,
        "Lasr Name " : req.body.last_name
    }

    console.log(response);
    resp.end(JSON.stringify(response));
})

var server = app.listen(3000,function(){
    var host = server.address().address;
    var port = server.address().port;

    console.log("应用实例，访问地址为 http://%s:%s", host, port);
})