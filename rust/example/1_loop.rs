fn main(){
    let mut a = 0;
    loop {
        println!("Hello rust!");
        if a == 10 {
            break;
        }
        a += 2;
    }
}