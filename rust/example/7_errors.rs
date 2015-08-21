/**
 * 可恢复的错误： 
    可恢复的错误是报告给用户的错误，用户可以重试该操作。 完全停止该过 程的可恢复错误并不严重。 
    它由Result <T，E>表示。 可恢复错误的示例是“找不到文件”。T＆E是通用参数。
        T->这是一种值，在成功的情况下返回一个’OK’变量。
        E->这是一种错误类型，在具有Err变体的故障情况下返回。
----------------------

不可恢复的错误： 
    当Rust报告一个不可恢复的错误时，那就是!宏停止执行程序。例如：“除以零”是不可恢复错误的示例
 */

enum Value {  
    Val,  
}  

fn get_number(_:Value)->i32 {   
   5  
}  
fn find_number(val:Value)-> &'static str {  
  match get_number(val) {  
    7 => "seven",  
    8=> "eight",  
    _=> unreachable!()  
    /**
     * unreachable! ：它用于无法访问的代码。 
     * 此宏很有用，因为编译器无法确定无法访问的代码。 
     * 在运行时由 unreachable! 执行。
     */
  }  
}  

fn main() {  
    let x : bool = false;  
    assert!(x==true); //断言失败则panic！！

    println!("{}", find_number(Value::Val));  
}