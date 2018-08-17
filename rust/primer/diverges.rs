fn main() {
    // let x = true;
    // let y: bool = false;

    // //# char
    // // 不像其它语言，这意味着 Rust 的 char 并不是 1 个字节，而是 4 个。
    // let x = 'x';
    let two_hearts = '💕';  //unicode
    println!("{}",two_hearts);
    diverges()
}

fn diverges() ->!{
    panic!("This function never returns!");
}
/*
 * panic! 是一个宏，类似我们已经见过的 println!() 。
 * 与 println!() 不同的是，panic!() 导致当前的执行线程崩溃并返回指定的信息。
 * 因为这个函数会崩溃，所以它不会返回，所以它拥有一个类型!，它代表“发散”。
 * 如果你想要更多信息，你可以设定 RUST_BACKTRACE 环境变量来获取 backtrace ：
 * RUST_BACKTRACE=0 ./diverges
 * RUST_BACKTRACE=1 cargo run
 */

// 关于注释 : /** a */ 会出错,'/*之后不能有连个星号'



