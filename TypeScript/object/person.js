// class person{
//     constructor(name:string){
//         this.name = name;
//     }
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    }
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
//     public name:string;
//     private age;
//     protected sex;
//     eat(){
//         console.info(this.name," -> eat()");
//     }
// }
var person = /** @class */ (function () {
    function person(name) {
        this.name = name;
        //需要注意。当在构造函数内部声明一个属性name的时候，需要指定访问控制符
    }
    person.prototype.eat = function () {
        console.log(this.name + "is eating lunch.");
    };
    return person;
}());
var employee = /** @class */ (function (_super) {
    __extends(employee, _super);
    function employee(name, code) {
        var _this = _super.call(this, name) ///子类的构造函数必须先用super关键字调用父类的构造函数，这是规定
         || this;
        // code = code //必须使用this.code 
        //果然微软还是那个微软,满满的洗佳佳的味道.
        _this.code = code;
        console.warn("New a emloyee object");
        return _this;
    }
    employee.prototype.work = function () {
        this.eat();
        this.doWork();
    };
    employee.prototype.doWork = function () {
        console.info(this.name, "is working,code : ", this.code);
    };
    return employee;
}(person));
//注意,ts在类型定义上与编译型语言类似,区分大小写
var Employee = /** @class */ (function (_super) {
    __extends(Employee, _super);
    function Employee(name, code) {
        var _this = _super.call(this, name) || this;
        _this.code = code;
        console.info("New an Employee object.");
        return _this;
    }
    Employee.prototype.work = function () {
        _super.prototype.eat.call(this);
        this.doWork();
    };
    Employee.prototype.doWork = function () {
        console.warn(this.name + "is using " + this.code
            + " for his job");
    };
    return Employee;
}(person));
var p1 = new person("Ethan");
p1.eat();
var e = new employee("EthanJay", "TypeScript");
e.work();
console.log("-----------------------------------------");
var em = new Employee("SergeyCheung", "golang");
em.work();
