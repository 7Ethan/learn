process.on('exit',function(code){
    setTimeout(() => {
        console.log('these code will never run...');
    }, 2);
    // 输出全局变量 __filename 的值
    console.log(__filename)
    // 输出全局变量 __dirname 的值
    console.log( __dirname);
    console.log('exit code is : ',code)
})

console.log("程序执行结束");