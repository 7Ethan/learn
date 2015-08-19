fn main(){
    for i in 1..11{
        println!("2 * {} = {}",i,i * 2);
    }

    let fruits=["mango","apple","banana","litchi","watermelon"]; 
    for i in fruits.iter(){
        print!("{},",i);
    }
    println!();
}