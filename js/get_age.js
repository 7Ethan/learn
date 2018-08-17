let obj = {
    birth:1999,
    get_age:function(){
        let b = this.birth;
        let fn = () => new Date().getFullYear() - b;
        return fn();
    }
};

let a = obj.get_age();

// let obj = {
//     birth:1999,
//     get_age:function(){
//         let b = this.birth;
//         let fn = () => new Date().getFullYear() - b;
//         return fn;
//     }
// };

// let a = obj.get_age()();
console.log("a is : ",a);