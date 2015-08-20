#[derive(Debug)]
enum Flagcolor {
    Orange,
    White,
    Red,
}

use Flagcolor::{Orange,White,Red};
// use Flagcolor::*;
/** 
 * 或者使用 use Flagcolor::*;
*/
fn main(){
    let _o = Orange;
    let _w = White;
    let _r = Red;

    println!("{:?}",_o);
    println!("{:?}",_w);
    println!("{:?}",_r);
}