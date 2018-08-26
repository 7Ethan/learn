import java.io._

class Person(name:String){
  println("Person info " + info())

  val age:Int = 18
  var hobby:String = "Baskelball"
  println("Person info : " + info())

  def info():String = {
    "name : %s,age : %d,hobby:%s".format(name,age,hobby)
  }
}

object Person{
  def main(args:Array[String]) {
    new Person("Ethan")
  }
}
