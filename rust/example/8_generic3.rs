 struct Program<T> {  
    a: T,  
    b: T,  
}  
impl<T> Program<T>   {  
    fn a(&self) -> &T {  
       &self.a  
    }  
}

fn main() {  
    let p = Program{ a: 5, b: 10 };  

    println!("p.a() is {}", p.a());  
}
