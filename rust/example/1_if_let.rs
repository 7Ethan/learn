fn main() {
    let a  = if false{
        100
    } else {
        // "hello" //ERROR!!!!
        99
    };
    println!("a value is :{}",a);
}