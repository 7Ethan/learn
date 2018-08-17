let aNumber:any = "hello"
console.log(aNumber)
console.log(aNumber = 10)
// 声明一个变量为任意值之后，对它的任何操作，返回的内容的类型都是任意值。
// 包括调用任何方法

//变量如果在声明的时候，未指定其类型，那么它会被识别为任意值类型
let something;
something = 'seven';
something = 7;

something.setName('Tom');