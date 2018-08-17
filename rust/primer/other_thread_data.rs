use std::thread;

fn main() {
    let data = vec![1,2,4,67,45,567,];
    let handle = thread::spawn(move|| {    //必须使用move,否则无法访问其他现成(这里是main线程)的数据
        // Rust 不知道这个新建线程会执行多久，所以无法知晓 data 的引用是否一直有效
        println!("{:?}",data);
    });

    // drop(data); //丢弃数据

    handle.join().unwrap();
}