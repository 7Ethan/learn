#[macro_use]
pub mod macros;

macro_rules! hello {
    () => (
        println!("Hello Ethan");
    )
}

fn main() {
    hello!();
    sayHello!();
}

