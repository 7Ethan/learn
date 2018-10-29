//util.inherits(constructor, superConstructor)是一个实现对象间原型继承 的函数。

var util = require('util');
function Base(){
    this.name = 'base';
    this.base = 999;
    this.sayHello = function (){
        console.log("Util's SAYHELLO functon");
    };
};

Base.prototype.showName = function(){
    console.log(this.name);
}

function Sub(){
    this.name = 'sub';
}

util.inherits(Sub,Base);
var objBase = new Base();
objBase.showName();
objBase.sayHello();
console.log(objBase);
console.log('---------------');

var objSub = new Sub();
// objSub.sayHello();
// Sub 仅仅继承了Base 在原型中定义的函数，
// 而构造函数内部创造的 base 属 性和 sayHello 函数都没有被 Sub 继承。
objSub.showName();
console.log(objSub);