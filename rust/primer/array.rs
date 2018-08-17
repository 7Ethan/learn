fn main() {
    let array = [0;20]; //创建一个容量为10,全为1的数组
    // println!("{}",a);
    for a in array.iter() {
        print!("{} ",a);
    }
    println!();

    range_number();
    map_method();
}

fn range_number(){
    for i in (0..10).rev(){
        print!("{} ",i);
    }
    /*
     * python 的写法是 :
     * for i in range(0,10,2):
     *      dosomething();
     */
    println!();
}

fn map_method(){
    for i in (1..10).map(|x| x * x ){
        print!("{} ",i);
    }
    println!();

}

// # 1.
// ----
// asdfa




