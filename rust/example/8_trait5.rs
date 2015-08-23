trait Sample  {  
    fn a(&self);  
    fn b(&self){  
        println!("Print b");  
    }   
}   

struct Example {  
    a:i32,  
    b:i32,  
}  

impl Sample for Example {  
  fn a(&self){  
    println!("Value of a is {}",self.a);  
  }  

  fn b(&self){  
    println!("Value of b is {}",self.b);  
  }  
}  

fn main() {  
  let r = Example{a:5,b:7};  
  r.a();  
  r.b();    
}
