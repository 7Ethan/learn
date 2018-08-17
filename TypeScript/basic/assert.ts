// let myFavoriteNumber = 'seven';
// myFavoriteNumber = 7;
let str = 'this is a string'
str = "a new string"

// index.ts(2,1): error TS2322: Type 'number' is not assignable to type 'string'.

// /**
//  * 一 . 会在没有明确的指定类型的时候推测出一个类型，这就是类型推论
//  * 二 . 如果定义的时候没有赋值，不管之后有没有赋值，都会被推断成 any 类型
//  * /

function GetLength(someThing:string|number):number{
    if ((<string>someThing).length){
        return (<string>someThing).length
    } else{
        return someThing.toString().length
    }
}

let len:number;
len = GetLength(12334)
console.log(len)
