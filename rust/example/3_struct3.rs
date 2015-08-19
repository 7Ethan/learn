#[derive(Debug)]
struct Square {
    a: f64,
}

impl Square {
    fn area(&self) -> f64{
        self.a * self.a 
    }
}

fn main() {
    let s =  Square{a:4.5};
    println!("The area of square is : {}",s.area());
}