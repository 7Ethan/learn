macro_rules! say_hello {
    () => {
        println!("{}", String::from("Hello Macro!"));
    };
}

fn main(){
    say_hello!();
}