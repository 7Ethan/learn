var express = require('express');
var app = express();
var fs = require('fs');

var bodyParser = require('body-parser');
var multer = require('multer');

app.use(express.static('public'));
// var urlencodeParser = bodyParser.urlencoded({extended:false})
app.use(bodyParser.urlencoded({extended:false}))

app.get('/index.html',function(req,resp){
    resp.sendFile(__dirname + '/' + 'upload_file.html');
})

var server = app.listen(3000,function(){
    var host = server.address().address,
    var port = server.address().port

    console.log("应用实例，访问地址为 http://%s:%s", host, port)
})