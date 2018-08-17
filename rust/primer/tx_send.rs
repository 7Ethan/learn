use std::thread;
use std::sync::mpsc;

fn main(){
    let (tx,rx) = mpsc::channel();

    thread::spawn(move || {
        let val = String::from("hi");
        tx.send(val).unwrap();
        // println!("val is {}",val); //发送之后不能在使用val这个值.所有权已经发生转移
    });

    let receiver = rx.recv().unwrap();
    println!("Got : {}",receiver);
}