trait A  {  
    fn f(&self);  
}  

trait B : A {  
    fn t(&self);  
}  

struct Example {  
    first : String,  
    second : String,  
}  

impl A for Example  {  
    fn f(&self) {  
        print!("{} ",self.first);  
    }  
 }  

 impl B for Example  {  
    fn t(&self) {  
        print!("{}",self.second);  
    }  
}  

fn main(){  
    let s = Example{first:String::from("Google"),second:String::from("rust-lang")};  
    s.f();  
    s.t();  
}
