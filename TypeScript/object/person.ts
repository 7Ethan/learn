// class person{
//     constructor(name:string){
//         this.name = name;
//     }

//     public name:string;
//     private age;
//     protected sex;

//     eat(){
//         console.info(this.name," -> eat()");
//     }
// }

class person{
    constructor(public name: string) {
    //需要注意。当在构造函数内部声明一个属性name的时候，需要指定访问控制符
    }

    private age;//只可以在类内部访问

    protected sex;//只可以在类内部或子类的内部被访问

    eat() {//如果没有访问控制符的话，默认为public
        console.log(this.name + "is eating lunch.")
    }
}

class employee extends person{
    code:string
    constructor(name:string,code:string){
        super(name) ///子类的构造函数必须先用super关键字调用父类的构造函数，这是规定
        // code = code //必须使用this.code 
        //果然微软还是那个微软,满满的洗佳佳的味道.
        this.code = code;
        console.warn("New a emloyee object")
    }

    work(){
        this.eat()
        this.doWork()
    }
    
    private doWork(){
        console.info(this.name,"is working,code : ",this.code)
    }
}

//注意,ts在类型定义上与编译型语言类似,区分大小写
class Employee extends person{
    constructor(name:string,code:string){
        super(name)
        this.code = code
        console.info("New an Employee object.")
    }
    code:string

    work(){
        super.eat()
        this.doWork()
    }

    private doWork(){
        console.warn(this.name + " is using " + this.code
                +" for his job")
    }
}

var p1 = new person("Ethan");
p1.eat();

var e = new employee("EthanJay","TypeScript")
e.work();

console.log("-----------------------------------------")

var em = new Employee("SergeyCheung","golang")
em.work()

