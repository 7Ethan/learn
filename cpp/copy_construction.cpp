/**
 * 拷贝构造函数是一种特殊的构造函数，它在创建对象时，是使用同一类中之前创建的对象来初始化新创建的对象。
 * 拷贝构造函数通常用于：
 * -> 通过使用另一个同类型的对象来初始化新创建的对象。
 * -> 复制对象把它作为参数传递给函数。防止原对象被修改(不希望被修改的情况下).
 * -> 复制对象，并从函数返回这个对象。防止不同对象指向同一内存
 * 
 */

#include <iostream>
 
using namespace std;
 
class Line{
   public:
      int getLength( void );
      string getName(void);
      Line();
      Line( int len,string name );             // 简单的构造函数
      Line( const Line &obj);      // 拷贝构造函数
      ~Line();                     // 析构函数
      //先释放堆,后释放栈.
      // 因为栈通常是一个程序运行的上下文,如果先释放栈,很可能堆中对象无法释放,发生内存泄漏
 
   private:
      int *ptr; 
      //成员变量有指针或者数组的时候要特别注意
      //要随时留意深拷贝与浅拷贝的区别.
      string name;

    /** 
     * 深拷贝：新建一个堆，然后就通过循环的方法把堆中的一个一个的数据拷过去。
     * 这样就可以避免在修改拷贝对象的数据时，同时改变了被拷贝对象的数据.
    */
   /**
    * 浅拷贝中，两个指向同一块内存。释放内存时会出现问题。
   */

  /**
   * 类包含指针对象时，先实例化指针对象a,b。再去调用父对象的构造函数。
   * 在销毁时也是先按照顺序销毁指针对象a,b；最后销毁父对象。
  */
};


Line::Line(){

}
 
// 成员函数定义，包括构造函数
Line::Line(int len,string name)
{
    this->name = name;
    cout << "调用构造函数 : Line(int len,string name) " << endl;
    // 为指针分配内存
    ptr = new int;
    *ptr = len;
}
 
Line::Line(const Line &obj){
    this->name = obj.name;
    cout << "调用拷贝构造函数并为指针 ptr 分配内存->" << "对象 name 属性的内存地址是:" << &this->name << endl;    
    // cout << "" << endl;
    ptr = new int;  //// 重点实现: 不是直接赋值，而是开辟了自己的内存地址
    // 这里实现的是深拷贝,避免不同对象指向同一内存地址(回收其中一个对象的内存后另外一个对象会执行空引用)
    *ptr = *obj.ptr; // 拷贝值
}
 
Line::~Line(void){
    cout << "释放内存" << endl;
    delete ptr;
}

int Line::getLength( void ){
    return *ptr;
}

string Line::getName(void){
    return name;
}
 
void display(Line obj){
   cout << "line 大小 : " << obj.getLength() <<endl;
    cout << "name is : " << obj.getName() <<endl;
}
 
// 程序的主函数
int main( ){
//    Line line1(10); 
    Line *line1 = new Line(10,"Ethan"); //这里调用的是带参数的构造函数
    cout << "A new object has been newed.." << endl;
 
   Line line2 = *line1; // !!!!!!!!这里两次调用拷贝构造函数
 
   display(*line1); //在这里编译器检测到line1对象已经没有使用,将其回收,调用析构函数.
   //并不会等到整个程序结束才会回收内存.
   
   
   display(line2);  // 对象作为参数传递,再一次调用拷贝构造函数!!!!!

     // ---test1 ---
    // Line *line4 = new Line(10);    //还是会调用构造函数.
    // display(*line4); //正常运行.

    // ---test2 ---
    // Line *line3 = new Line();    //还是会调用构造函数.
    // display(*line3);     //可以编译,运行出错:8613 segmentation fault (core dumped)
    // //原因是访问了没有初始化的内存.
   return 0;
}
