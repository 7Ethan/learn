use std::thread;
use std::sync::mpsc;
use std::time::Duration;

fn main(){
    let (tx,rx) = mpsc::channel();

    //let tx1 = mpsc::Sender::clone(&tx);
    //两种写法都可以
    let tx1 = tx.clone();

    thread::spawn(move || {
        let vals = vec![
            String::from("你好"),
            String::from("这里是"),
            String::from("rust世界的"),
            String::from("新手园地"),
        ];

        for val in vals {
            tx.send(val).unwrap();
            // thread::sleep(Duration::new(3,0));
            thread::sleep(Duration::from_millis(500));
        }
    });


    thread::spawn(move || {
        let vals = vec![
            String::from("AA"),
            String::from("BB"),
            String::from("CC"),
            String::from("DD"),
            String::from("RUST!!!!"),
        ];

        for val in vals {
            tx1.send(val).unwrap();
            // thread::sleep(Duration::new(2,0));
            thread::sleep(Duration::from_millis(800));
        }
    });

    for receiver in rx{
        println!("Receive data from spawn thead :{}",receiver);
    }
}