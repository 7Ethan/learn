def whileloop: Unit ={
  var i = 0
  while (i <= 3){
    println(i)
    i += 1
  }
}

def forloop: Unit ={
  println("for loop using java-style iteration.")
  for ( i <- until args.length){
    println(args[i])
  }
}

whileloop

forloop