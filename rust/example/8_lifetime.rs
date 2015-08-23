#[derive(Debug)]
struct Example<'a> {
    x: & 'a i32
}

impl <'a> Example<'a> {
    fn display(&self){
        println!("Value is  : {}",self.x);
    }
}

fn main(){
    let y = &999;
    let e = Example{x:y};
    e.display();
}