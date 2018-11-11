package hello

class greeter(val name:String){
    fun greet(){
        println("Hello $name")
    }
}

fun main(args: Array<String>) {
    greeter("Kotlin").greet()
}