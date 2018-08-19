class person{
    constructor(public name: string) {
    //需要注意。当在构造函数内部声明一个属性name的时候，需要指定访问控制符
        console.warn("New an person object")
    }

    private age;//只可以在类内部访问

    protected sex;//只可以在类内部或子类的内部被访问

    eat() {//如果没有访问控制符的话，默认为public
        console.log(this.name + "is eating lunch.")
    }
}

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

// var workers:Array<person> = []

// workers[0] = new person("王五")
// workers[0].eat()
// workers[1] = new Employee("小张","java  ")
// // workers[1]

var workers:Array<Employee> = []

// workers[0] = new person("王五") // Error
// workers[0].eat()
// workers[1] = new Employee("小张","java  ")
// workers[1]