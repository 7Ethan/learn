use std::thread;
use std::sync::mpsc;
use std::time::Duration;

fn main(){
    let (tx,rx) = mpsc::channel();  //tx,rx 顺序不能乱!! 

    thread::spawn(move || {
        let vals = vec![
            String::from("你好"),
            String::from("这里是"),
            String::from("rust世界的"),
            String::from("新手园地"),
        ];

        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::new(1,0));
        }
    });

    for receiver in rx {
        println!("Recive from spwan thread : {}",receiver);
    }
}