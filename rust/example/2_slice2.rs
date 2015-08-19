fn main()  

    let arr = [100,200,300,400,500,600];  
    let mut i=0;  
    let a=&arr[1..=3];  
    let len=a.len();  
    println!("Elements of 'a' array:");  
    while i<len  
    {  
     println!("{}",a[i]);  
     i=i+1;  
    }  
}
