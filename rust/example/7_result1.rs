/**
 *  可恢复的错误是那些完全停止程序并不严重的错误。可以处理的错误称为可恢复错误。
    它由Result <T，E>表示。 结果<T，E>是枚举，由两个变体组成，即OK <T>和Err <E>，它描述了可能的错误。
        OK <T>：’T’是一种值，它在成功情况下时返回OK变量，这是一个预期的结果。
        Err <E>：’E’是一种错误，它在失败情况下时返回ERR变量，这是意想不到的结果。
 */
use std::fs::File;  
fn main() {  
    // let f:u32 = File::open("vector.txt"); 

    let f = File::open("vector.txt");  
    match f {  
        Ok(file) => file,  
        Err(error) => {  
            panic!("There was a problem opening the file: {:?}", error)  
        },  
    };  

    File::open("hello.txt").unwrap(); 
    /**
     * 结果<T，E>有许多方法可以提供各种任务。 其中一种方法是unwrap()方法。 
     * unwrap()方法是匹配表达式的快捷方法。 unwrap()方法和匹配表达式的工作方式是一样的。
        --如果Result值是OK变量，则unwrap()方法返回OK变量的值。
        --如果Result值是Err变量，那么unwrap()方法会调用panic!宏。
     */

    File::open("hello.txt").expect("Not able to find the file hello.txt");  
    /**
     * expect()方法的行为方式与unwrap()方法相同，即两种方法都会引起panic!宏显示错误信息。
     * expect()和unwrap()方法之间的区别在于，
     *      错误消息作为参数传递给expect()方法，而unwrap()方法不包含任何参数。  
     *      因此，可以说expect()方法可以跟踪panic!的来源更容易。
     */
}

