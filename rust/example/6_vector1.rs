fn value(number:Option<&f64>){
    match number{
        Some(number) => println!("Fourth element of a vector is {}",number),
        None => println!("None"),
    }
}

fn main(){
    let v = vec![20.0,30.0,40.0,50.0]; 
    let a : Option<&f64> = v.get(3);
    value(a);

    // v.get（） 方法好处是当当访问不存在的数据时，会返回None，而不会panic
    /**
     * 编译期间并不会发现诸如数组越界等问题！！
     *  println!("{}",v[100]);  
     * //thread 'main' panicked at 'index out of bounds: the len is 4 but the index is 100'
     */
    let b : Option<&f64> = v.get(100);
    value(b)
}