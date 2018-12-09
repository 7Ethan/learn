macro_rules! create_function {
    ($func_name:ident) => {
        fn $func_name(){
            println!("You are calling function : {:?}()", stringify!($func_name));
        }
    };
}

create_function!(func1);
create_function!(func2);

macro_rules! print_result {
    ($express:expr) => {
        println!("{:?} = {:?} ", stringify!($express),$express);
    };
}

fn main(){
    func1();
    func2();
    print_result!(1u32 + 1);

    print_result!({
        let x = 2u32;
        x * x   
    })
}