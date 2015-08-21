/**
 * String
    字符串被编码为UTF-8序列。
    在堆内存上分配一个字符串。
    字符串的大小可以增长。
    它不是以空(null)值终止的序列。
    ----------------------

&str
    &str也称为字符串切片。
    它由&[u8]表示，指向UTP-8序列。
    ＆str用于查看字符串中的数据。
    它的大小是固定的，即它不能调整大小。
    -----------------------

String 和 &str 的区别
    String是一个可变引用，而&str是对该字符串的不可变引用，即可以更改String的数据，但是不能操作&str的数据。
    String包含其数据的所有权，而&str没有所有权，它从另一个变量借用它。
 */

fn main() {  
    let mut s=String::from("java is a");  
    s.push_str(" programming language");  
    print!("{}",s);  

    let s1 = String::from("C");  
    let s2 = String::from("is");  
    let s3 = String::from("a");  
    let s4 = String::from("programming");  
    let s5 = String::from("language.");  
    let s = format!("{} {} {} {} {}",s1,s2,s3,s4,s5);  
    print!("{}",s);
}