fn main() {  
    let a = Some(5);  
    match a {  
        Some(5) => println!("five"),  
        _ => (), 
    }

    let a=Some(3);  
    if let Some(3)=a{  
        println!("three");  
    }
}