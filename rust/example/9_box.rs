fn main(){
    let a = Box::new(1);
    println!("value of a is : {}",a);
    /**
     * Box <T>是一个智能指针，指向在类型为T的堆上分配的数据。
     * Box <T>允许将数据存储在堆而不是堆栈上。
     */
}