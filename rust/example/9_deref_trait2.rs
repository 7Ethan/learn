/**
 * Deref强制是将实现Deref trait的引用转换为Deref可以将原始类型转换为的引用的过程。
 * Deref强制是对函数和方法的参数执行的。
 * 当将特定类型的引用传递给与函数定义中的参数类型不匹配的函数时，Deref强制自动发生。
 */

struct MyBox<T>(T);  

use :: std::ops::Deref;  

impl<T> MyBox<T>  {  
    fn hello(x:T)->MyBox<T>  {  
        MyBox(x)  
    }  
}  

impl<T> Deref for MyBox<T> {  
    type Target = T;  
    fn deref(&self) ->&T{  
        &self.0
    }  
}  

fn print(m : &i32) {  
    print!("{}",m);  
}  

fn main() {
    let b = MyBox::hello(5);  
    print(&b);  
}