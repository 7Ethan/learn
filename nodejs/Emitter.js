// 引入 events 模块
var events = require('events');
var hello = require('./web/HelloServer.js').world;
hello();

var emitter = new events.EventEmitter();
const data = 'data_received';

// 创建事件处理程序
var connectHandler = function connected(){
    console.log('connect suceessfully!');
    emitter.emit('data_received');
}

// 绑定 connection 事件处理程序
emitter.on('connection',connectHandler)

// 使用匿名函数绑定 data_received 事件
emitter.on(data,function (){
    console.log('received data');
})

// 触发 connection 事件
emitter.emit('connection');

console.log('Program is over.');
