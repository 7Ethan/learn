/**
 * drop trait用于在值超出范围时释放文件或网络连接等资源。
 * drop trait用于释放Box <T>指向的堆上的空间。
 * drop trait用于实现drop()方法，该方法对self进行可变引用
 */

struct Example{
    a:i32
}

impl Drop for Example{
    /**
     * 相当于 C++ 中手动实现析构函数，与 RALL 技术有相似的地方。
     */
    fn drop(&mut self){
        println!("Dropping the instance of Example with data : {}",self.a);
    }
}

fn main(){
    let e = Example{a : 100};
    drop(e);    //手动调用，而不是等到程序执行完毕再回收内存！！！！！！！！！！！！！！！！！！！
    let e2 = Example{a : 18};
    println!("Created 2 instance !" );
}