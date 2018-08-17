//引用规则:
/**
 * 1. 在任意给定时间，只能 拥有如下中的一个：
    -> 一个可变引用。
    -> 任意数量的不可变引用。
  2. 引用必须总是有效的。
 */


fn main() {
    let s = String::from("Hesdfsagaaq43llo");
    let mut str = String::from("Hello");    
    println!("length of string is :{}",caculate_length(&s) );
    change(&mut str);
    println!("{}",str );
}

//& 符号就是 引用，他们允许你使用值但不获取其所有权
fn caculate_length(str:&String) ->usize{
    str.len()
}

fn change(str:&mut String){
    str.push_str(",aaaaaaaa");
}

fn borrow(){
    //不能在拥有不可变引用的同时拥有可变引用
    let mut s = String::from("hello");
    /**
     * 多个不可变引用是没有问题的因为没有哪个只能读取
     * 数据的人有能力影响其他人读取到的数据
     */
    let r1 = &s; // no problem
    let r2 = &s; // no problem
    // let r3 = &mut s; // BIG PROBLEM
}