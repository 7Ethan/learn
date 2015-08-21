use std::io;  
use std::io::Read;  
use std::fs::File;  

fn main() {  
    let a = read_username_from_file();  
    print!("{:?}",a);  
}  

fn read_username_from_file() -> Result<String, io::Error> {  
    let mut s = String::new();  
    File::open("a.txt")?.read_to_string(&mut s)?;  
    /**
     * 将read_to_string()的调用链接到 File::open("a.txt")? 的调用结果。
     * 如果两个函数(即 read_to_string() 和 File::open("a.txt") 成功，则返回OK值，否则返回错误值。
     */
    Ok(s)  
}
