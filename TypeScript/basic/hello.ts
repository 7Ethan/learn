function sayHello(person:string){   
    //TypeScript 只会进行静态检查，如果发现有错误，编译的时候就会报错。
    //Ts编译也会讲注释传给js文件
    return "Hello  " + person
}

let person = "Ethan";
console.log(sayHello(person));