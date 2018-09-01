use std::sync::{Arc,Mutex};
use std::thread;

fn main(){
    let data = Arc::new(Mutex::new(vec![1u32,3,2]));

    for i in 0..3{
        let mut data = data.clone();
        thread::spawn(move ||{
            let mut data = data.lock().unwrap();
            data[i] += 1;
        });
    }

    thread::sleep_ms(50);
}
