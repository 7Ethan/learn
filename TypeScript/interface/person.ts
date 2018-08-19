interface iPerson{
    name:string
    age: number
}

class person{
    constructor(public cinfig:iPerson){

    }
    work(){
        console.log("Person implement iPerson ,working...")
    }
}

var p1 = new person({
    name:"Ethan",
    age:18,
})