/**
 * Deref <T> trait用于自定义解除引用运算符(*)的行为。
 */
#[derive(Debug)]
struct MyBox<T> {
    field: T
}

use std::ops::Deref;
impl<T> Deref for MyBox<T>{
     /**
     * 用处其实就是解引用，而不是使用 * 进行解引用
     */
    type Target = T;
    fn deref(&self) -> &T{
        &self.field
    }
}


fn main(){
    let b = MyBox{field : String::from("Deref trait!")};
    println!("value of mybox is : {}",(b.deref()));
    /**
     * 通过使用MyBox类型的实例b.deref()调用deref()方法，
     * 然后取消引用从deref()方法返回的引用。
     */
}