function makecounter() {
    var i = 0;
    return function(){
        console.log(++i);
    }
}

f = makecounter();
f();    