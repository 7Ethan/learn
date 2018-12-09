macro_rules! find_min {
    ($x:expr) => {$x};
    ($x:expr,$($y:expr),+)=>{
        std::cmp::min($x,find_min!($($y),+));
    }
}

fn main(){
    println!("{}",find_min!(1u32));
    println!("{}",find_min!(32u32,56u32));
    println!("{}",find_min!(4u32,2 * 1u32,234u32));
}