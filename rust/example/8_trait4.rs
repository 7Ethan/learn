/**
 * 使用where子句。
    - 使用出现在括号{之前的where子句来编写绑定。
    - where子句也可以应用于任意类型。
    - 当使用where子句时，它使语法比普通语法更具表现力。
 */

trait Perimeter  {  
    fn a(&self)->f64;  
}  

struct Square {  
    side : f64,  
}

impl Perimeter for Square {  
    fn a(&self)->f64 {  
        4.0*self.side  
    }  
}  

struct Rectangle {  
    length : f64,  
    breadth : f64,  
}  

impl Perimeter for Rectangle {  
    fn a(&self)->f64 {  
        2.0*(self.length+self.breadth)  
    }  
}  
fn print_perimeter<Square,Rectangle>(s:Square,r:Rectangle)  
  where Square : Perimeter,  
        Rectangle : Perimeter  
        {  
          let r1 = s.a();  
          let r2 = r.a();  
          println!("Perimeter of a square is {}",r1);  
          println!("Perimeter of a rectangle is {}",r2);  
        }  


fn main() {  
    let sq = Square{side : 6.2};  
    let rect = Rectangle{length : 3.2,breadth:5.6};  
    print_perimeter(sq,rect);  
}