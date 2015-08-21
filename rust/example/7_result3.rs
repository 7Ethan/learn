/**
 * 使用?运算符减少了代码的长度。
 * ?运算符是匹配表达式的替换意味着?运算符的工作方式与匹配表达式的工作方式相同。
 */

use std::io;  
use std::io::Read;  
use std::fs::File; 

fn main() {  
  let a = read_username_from_file();  
  print!("{:?}",a);  
}  

fn read_username_from_file() -> Result<String, io::Error> {  
    let mut f = File::open("a.txt")?;  
    let mut s = String::new();  
    f.read_to_string(&mut s)?;  
    /**
     * ?运算符是在Result值类型之前使用。 
     * 如果Result为OK，则返回OK变体的值，
     * 如果Result为Err，则返回错误信息。
     */
    Ok(s)  
}