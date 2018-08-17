use std::fmt;

// #[derive(Debug)]
struct User {
    name: String,
    email:String,
    age:i32
}

impl fmt::Display for User {
    fn fmt(&self,f:&mut fmt::Formatter) ->fmt::Result{
        write!(f, "姓名是:{},年龄:{},邮箱:{}",self.name,self.age,self.email)
    }
}

fn main() {
    let user = User{name:String::from("Ethan"),email:String::from("Ethan@123.com"),age:19};
    // println!("{}", u)
    // println!("user is :{:#?}",user);
    println!("{}",user);
    print_number();
}

fn print_number(){
    let arr = [1,234,65,78,345,98];
    print!("{:#?}",arr);
    println!();
}