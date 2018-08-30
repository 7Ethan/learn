use std::thread;
use std::sync::mpsc;
const THREAD_COUNT: i32 = 2;

fn main() {
    let (tx,rx):(mpsc::Sender<i32>,mpsc::Receiver<i32>) = 
        mpsc::channel();

    for id in 0..THREAD_COUNT{
        let thread_tx = tx.clone();
        thread::spawn(move || {
            thread_tx.send(99 + id).unwrap();
        });
    }

    // thread::sleep_ms(1000);
    std::thread::sleep_ms(1000);
    println!("Wake up");

    for _i in 0..THREAD_COUNT{
        println!("{}", rx.recv().unwrap());
    }
    // thread::sleep_ms(1000);
    std::thread::sleep_ms(1000);
}