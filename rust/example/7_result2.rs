use std::io;  
use std::io::Read;  
use std::fs::File; 

fn main() {  
    let a = read_username_from_file();  
    print!("{:?}",a);  
}  

/**
 * read_username_from_file()函数返回Result <T，E>类型的值，
 * 其中'T'是String类型，'E'是io类型：Error。
 * -------
 * 如果函数成功，则返回一个包含String的OK值，如果函数失败，则返回Err值。
 */
fn read_username_from_file() -> Result<String, io::Error> {  
    let f = File::open("a.txt");

    let mut f = match f {  
        Ok(file) => file,  
        Err(e) => return Err(e),  
    };  

    let mut s = String::new();  
    match f.read_to_string(&mut s) {  
        Ok(_) => Ok(s),  
        Err(e) => Err(e),  
    }  
}