use std::fmt;
use std::sync::mpsc;
use std::thread;
use std::rc::Rc;

pub struct Student {
    id: u32
}

impl fmt::Display for Student {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "student {}", self.id)
    }
}

fn main() {
    /**
     * error[E0277]: the trait bound `std::rc::Rc<Student>: std::marker::Send` is not satisfied
  --> channel_struct.rs:22:5
   |
22 |     thread::spawn(move || {
   |     ^^^^^^^^^^^^^ `std::rc::Rc<Student>` cannot be sent between threads safely
   |
   = help: the trait `std::marker::Send` is not implemented for `std::rc::Rc<Student>`
   = note: required because of the requirements on the impl of `std::marker::Send` for `std::sync::mpsc::Sender<std::rc::Rc<Student>>`
   = note: required because it appears within the type `[closure@channel_struct.rs:22:19: 27:6 tx:std::sync::mpsc::Sender<std::rc::Rc<Student>>]`
   = note: required by `std::thread::spawn`

error: aborting due to previous error

For more information about this error, try `rustc --explain E0277`.

     */
    // 创建一个通道
    let (tx, rx): (mpsc::Sender<Rc<Student>>, mpsc::Receiver<Rc<Student>>) = 
        mpsc::channel();

    // 创建线程用于发送消息
    thread::spawn(move || {
        // 发送一个消息，此处是数字id
        tx.send(Rc::new(Student{
            id: 1,
        })).unwrap();
    });

    // 在主线程中接收子线程发送的消息并输出
    println!("receive {}", rx.recv().unwrap());
}