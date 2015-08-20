mod outer{
    pub fn a() {  
     println!("function a");          
    }  
    fn b() {  
        println!("function b");  
    }  

    mod inner{
        pub fn c(){
            println!("function c"); 
        }

        fn d(){
            println!("function d"); 
        }
    }
}

fn main(){
    outer::a();
    // outer::b(); //error[E0603]: function `b` is private
    // outer::inner::c(); //error[E0603]: module `inner` is private
    // outer::inner::d(); //error[E0603]: module `inner` is private
}

