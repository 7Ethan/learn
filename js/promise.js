function test(resolve,reject) {
    var timeout = Math.random() * 2;
    console.log('set timeout to: ' + timeout + ' seconds.');
    setTimeout(function(){
        if (timeout < 1){
            console.log("can resolve()...")
            resolve('200 OK')
        } else{
            console.log("Can't resolve ...")
            reject('timeout in ' + timeout + ' seconds.');
        }
    },timeout * 1000)
}

new Promise(test).then(function (result) {
    console.log('Success : ' + result)
}).catch(function (reason) {
    console.log('失败：' + reason);
})

// var p1 = new Promise(test)
// var p2 = p1.then(function(result){
    
// })
// var p3 = p2.catch(function (reason) {
    
// });