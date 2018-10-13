use std::string::String;

// fn foo<'a,'b>(x:&'a str,y:&'b str) ->Option<&'b str>{
//     if x.len() > y.len() {
//         Some(x)
//     } else{
//         Some(y)
//     }
// }

// fn main(){
//     // let s1 = String::from("Hola");
//     // let s2 = String::from("world");
//     // println!("{}",foo(s1.as_str(), s2.as_str()));
//     //------------------------------------------------

//     let s1 = String::from("你好");
//     let result;
//     {
//         let s2 = String::from("world");
//         result = foo(&s1, &s2);
//     }   
//     println!("{:?}", result);
// }


// fn search<'a, 'b>(needle: &'a str, haystack: &'b str) -> Option<&'b str> {
//     // imagine some clever algorithm here
//     // that returns a slice of the original string
//     let len = needle.len();
//     if haystack.chars().nth(0) == needle.chars().nth(0) {
//         Some(&haystack[..len])
//     } else if haystack.chars().nth(1) == needle.chars().nth(0) {
//         Some(&haystack[1..len+1])
//     } else {
//         None
//     }
// }

// fn main() {
//     let haystack = "hello little girl";
//     let res;
//     {
//         let needle = String::from("ello");
//         res = search(&needle, haystack);
//     }
//     println!("{:?}",res);
//     match res {
//         Some(x) => println!("found {}", x),
//         None => println!("nothing found")
//     }
//     // outputs "found ello"
// }

