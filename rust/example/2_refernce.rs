fn main(){
    let str = String::from("Hello rust");
    let l = change_str(&str);
    println!("str's length is : {}",l);
}

fn change_str(s:&String) -> usize {
    s.len()
}