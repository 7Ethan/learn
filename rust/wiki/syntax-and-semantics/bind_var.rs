

#[derive(Debug)]
struct Person {
    name: Option<String>
}


fn main() {
    let name = "Ethan".to_string();
    let x:Option<Person> = Some(Person{name:Some(name)});

    match x {
        Some(Person{name:ref a @Some(_),..}) => println!("{:?}", a),
        _ => println!(">>:")
     }
}


