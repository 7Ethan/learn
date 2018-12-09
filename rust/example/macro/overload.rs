macro_rules! test {
    ($left:expr;and $right:expr) => {
        println!("{:?} and {:?} is {:?}",
                stringify!($left),
                stringify!($right),
                $left && $right);
    };
    ($left:expr; or $right:expr) =>{
        println!("{:?}{:?}{:?}",
                stringify!($left),
                stringify!($right),
                $left || $right);
    };
}

fn main() {
    test!(2i32 + 1 == 3i32; and 2i32 * 2 == 4i32);
    test!(true;and false);
    test!(true;or false);
}