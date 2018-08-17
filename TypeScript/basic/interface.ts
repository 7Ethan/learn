interface Person{
    readonly id:number; //只读属性
    name:string;
    age?:number;
    // [propName:string]:string; //number不是string的子属性,会报错.
    [propName:string]:any;
}

let Ethan:Person = {
    id:1,
    name:"Ethan",
    age:25,
    gender:'male'   
};


//一旦定义了任意属性，那么确定属性和可选属性都必须是它的子属性