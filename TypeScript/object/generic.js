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
var person = /** @class */ (function () {
    function person(name) {
        this.name = name;
        //需要注意。当在构造函数内部声明一个属性name的时候，需要指定访问控制符
        console.warn("New an person object");
    }
    person.prototype.eat = function () {
        console.log(this.name + "is eating lunch.");
    };
    return person;
}());
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
        console.warn(this.name + " is using " + this.code
            + " for his job");
    };
    return Employee;
}(person));
var workers = [];
workers[0] = new person("王五");
workers[1] = new Employee("小张", "java  ");
