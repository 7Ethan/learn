#[derive(Debug)]
struct User {
    email: String,
    username:String,
    active:bool,
    sign_count:i64,
}

fn main() {
    let mut user1 = User{
        email:String::from("sergey@123.com"),
        username:String::from("SergeyChang"),
        active:true,
        sign_count:1,
    };
    user1.email = String::from("Ethan@123.com");
    println!("{}",user1.email );
}

fn printUser(user: User) {
    unimplemented!();
}

fn bind_user(email:String,username:String)-> User{
    User{
        username,
        email,
        active: true,
        sign_count: 1,
    }
}