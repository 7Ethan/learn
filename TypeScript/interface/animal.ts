interface Animal {
    eat()
}

class dog implements Animal{
    eat(){
        console.log("Dog has implement the interface...")
    }
}

class Tiger implements Animal{//Tiger 类实现了Animal接口
    eat() {
        console.log("i eat meat")
    }
}

