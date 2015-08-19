fn main(){
    // let str = "Hello Yiibai" ;    不能使用这样的变量进行切片
    //str的类型是&str，它是指向二进制文件特定点的切片。字符串文字是不可变的

    let str = String::from("google.com tutorial");
    let google = &str[0..9];
    let t1 = &str[11..=17];
    let t2 = &str[11..];
    println!(" --> {} , ==> {} ,t2 := {}",google,t1,t2);
    //输出结果：--> google.co , ==> tutoria ,t2 := tutorial
}