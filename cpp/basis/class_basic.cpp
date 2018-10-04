#include <iostream>
 
using namespace std;
 
class Line
{
   public:
      void setLength( double len );
      double getLength( void );
      Line();   // 构造函数声明
      Line(double);
      ~Line();  // 析构函数声明

    protected:
        bool isOK;
        //子类可以通过成员函数;来访问,但是不能直接访问.
        //保护成员在派生类（即子类）中是可访问的.
 
   private:
      double length;
      //私有成员变量或函数在类的外部是不可访问的，
      //甚至是不可查看的。只有类和友元函数可以访问私有成员
};

class point:Line{
    public:
        double x;
        double y;
};
 
// 成员函数定义，包括构造函数
Line::Line(void){
    cout << "Object is being created" << endl;
}

Line::Line(double len):length(len){
    cout << "Object is being created, length = " << len << endl;
}

Line::~Line(void){
    cout << "Object is being deleted" << endl;
}
 
void Line::setLength( double len ){
    length = len;
}
 
double Line::getLength( void ){
    return length;
}
// 程序的主函数
int main( ){
//    Line line(100.1);
 
//    // 设置长度
//    line.setLength(6.0); 
//    cout << "Length of line : " << line.getLength() <<endl;


    Line *line = new Line(99.9);    //new返回的是一个指针地址,所以用   -> 去引用.
    if (NULL == line){
        cout << "error accur when new a object!!!" << endl;
        return 0;
    }
    line->setLength(1.1);

    cout << "Length of line : " << line->getLength() <<endl;
    
    delete line;    //堆上分配的内存必须手动删除
    //否则析构函数不会执行,堆上分配的内存不会被释放!!!!!!!
    line = NULL;
    //析构函数在main退出前执行.

    cout << "main routine end"<< endl;

    //------test ----
    // point *p = new point();
    // p->x = 1.0;
    

   return 0;
}