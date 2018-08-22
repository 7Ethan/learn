fn main() {
    // let add_one = |x| x + 1;
    // let add_one = |x:i32| x +1;
    let add_one = |x:i32| ->i32 {x+1};
    println!("{}",add_one(1) );
}