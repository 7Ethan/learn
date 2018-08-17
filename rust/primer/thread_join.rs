use std::thread;

fn main(){
    let handler = thread::spawn(move || {
        for i in 1..100 {
            println!("Number {} from spawned thread",i );
        }
    });


    for i in 1..50{
        println!("Main Thread :{}", i);
    }

    handler.join();

}