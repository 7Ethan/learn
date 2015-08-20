fn check_even(number:i32) -> Option<i32>{
    /**
     * 选项 <T>包含两个变体：
            None：表示失败或缺少值。
            一些(值)：它是一个用T包装值的元组结构。
     *  */
    Some(number%2)
}

fn even_number(n:i32){
    let num = n;
    match check_even(num) {
        None => println!("None!"),
        Some(n) => {
            if n == 0{
                println!("{} is a even number .",num);
            } else{
                println!("{} is  a odd number . ",num);
            }
        },
    }
}

fn main(){
    even_number(2);
    even_number(3);
}