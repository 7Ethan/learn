fn add_one(x: &mut i32) {
    *x += 1;
}

fn main(){
    let mut x : i32 = 99;
    add_one(&mut x);
    println!("Now x is : {}", x);

    // ERROR!!!! 
    // let mut str=String::from("Yiibai");  
    // let a= &mut str;  
    // let b= &mut str;
}