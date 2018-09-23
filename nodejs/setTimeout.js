var emitter = require('events').EventEmitter;
var event = new emitter();

event.on('new_event',function(){
    console.log('that is a new event');
})

setTimeout(function(){
    event.emit('new_event');
},1000)