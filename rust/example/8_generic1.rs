 fn main()  
{  
  let a = vec![1,2,3,4,5];  
  let b = vec![2.3,3.3,4.3,5.3];  
  let result = add(&a);  
  let result1 = add(&b);  
  println!("The value of result is {}",result);  
  println!("The value of result1 is {}",result1);  
}  

fn add<T>(list:&[T])->T {  
  let mut c =0;  
  for &item in list.iter(){  
    c= c+item;  
  }  
}