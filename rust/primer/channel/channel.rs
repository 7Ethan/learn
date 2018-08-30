use std::thread;
use std::sync::mpsc;

fn main() {
    let (tx,rx):(mpsc::Sender<i32>,mpsc::Receiver<i32>) = 
        mpsc::channel();
    thread::spawn(move || {
        tx.send(1).unwrap();
    });

    println!("{}", rx.recv().unwrap());
}

