use std::fmt::{Debug, Display};  

fn compare_prints<T: Debug + Display>(t: &T) {  
    /**
     * Display和Debug特性通过使用 + 运算符限制为类型 T
     */
    println!("Debug: '{:?}'", t);  
    println!("Display: '{}'", t);  
}  

fn main() {  
    let string = "7Ethan";  
    compare_prints(&string);  
}