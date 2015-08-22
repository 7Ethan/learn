struct Triangle {  
  base : f64,  
  height : f64,  
}  

trait HasArea {  
  fn area(&self)->f64;  
}  

impl HasArea for Triangle  {  
  fn area(&self)->f64 {  
    0.5*(self.base*self.height)  
  }  
}  
fn main() {  
    let a = Triangle{base:10.5,height:17.4};  
    let triangle_area = a.area();  
    println!("Area of a triangle is {}",triangle_area);   
}