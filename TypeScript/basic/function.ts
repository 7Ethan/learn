// 函数声明（Function Declaration）
function sum(x, y) {
    return x + y;
}

// 函数表达式（Function Expression）
let mySum = function (x, y) {
    return x + y;
};  


//----------------------
interface SearchFunc{
    (source:string,subString:string):boolean
}

let mySearchFunc:SearchFunc;
mySearchFunc = function(a:string,b:string){
    return a.search(b) !== -1;
}
console.log(mySearchFunc("1234","12"))


// -----------------------
function buildName(firstName: string, lastName?: string) {
    if (lastName) {
        return firstName + ' ' + lastName;
    } else {
        return firstName;
    }
}
let tomcat = buildName('Tom', 'Cat');
let tom = buildName('Tom');
/**
 * 可选参数后面不允许再出现必须参数了!!!!!!!!!!!!!!!!!!!
 */

function twosum (x:number,y:number):number{
    return x + y
}


