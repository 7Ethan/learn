fn main(){
    let v =vec![20,30,40,50];  //不需要手动释放。离开作用域会被自动清理。
    println!("Elements of vector are :");  
    for iter in v {
        println!("{}",iter);
    }

    let mut v =vec![20,30,40,50];  
    print!("Elements of vector are :");  
    for i in &mut v {  
        *i+=20;  
        print!("{} ",i);  
    }

    let mut v=Vec::new();  
    v.push('j');  
    v.push('a');  
    v.push('v');  
    v.push('a');  
    for i in v {  
        print!("{}",i);  
    }  
}