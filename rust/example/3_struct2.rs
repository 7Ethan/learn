#[derive(Debug)]
struct Rect {
    width: f64,
    hight:f64,
}

fn area(r : &Rect) ->f64{
    r.width * r.hight
}

fn main(){
    let rect = Rect{width:3.0,hight:4.0};
    println!("Rect's area is : {}",area(&rect));
}
