enum Computerlanguage  {  
    C,  
    Cplus,  
    Java,  
    Csharp,  
}  

fn language(language:Computerlanguage) {  
    match language {  
        Computerlanguage::C=> println!("C language"),  
        Computerlanguage::Cplus=> println!("C++ language"),  
        Computerlanguage::Java=> println!("Java language"),  
        Computerlanguage::Csharp=> println!("C# language"),  
    }  
}  

fn main() {  
    language(Computerlanguage::C);  
    language(Computerlanguage::Cplus);  
    language(Computerlanguage::Java);  
    language(Computerlanguage::Csharp);  
}