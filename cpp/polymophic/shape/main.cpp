#include <iostream>
#include "circle.h"
#include "rect.h"
#include <stdlib.h>

using namespace std;

int main(int argc, char const *argv[])
{
    /**
     * 一句话:多态->允许将子类类型的指针赋值给父类类型的指针
     * 多态指同一个实体同时具有多种形式: 比方说一只动物,它能够调用狗叫的方法.
     *  c++中多态是通过虚函数来实现的.  
    */
   /**
    * C++中，实现多态有以下方法：
    *       虚函数，抽象类，覆盖，模板（重载和多态无关）。
   */
    shape *s1 = new circle(2.0);
    shape *s2 = new rect(3.0,4.0);

    s1->calcArea();
    s2->calcArea();

    delete s1;
    s1 = NULL;
    delete s2;
    s2 = NULL;

    system("pause");
    return 0;
}


// #include <iostream>
// #include "circle.h"
// #include "rect.h"
// #include <stdlib.h>
// using namespace std;

// int main()
// {
//     // 定义两个父类指针，指向子类对象
//     shape *shape1 = new circle(3.0);
//     shape *shape2 = new rect(3.0, 4.0);

//     shape1->calcArea();
//     shape2->calcArea();
//     //当基类不添加virtual时。打印两遍基类的。

//     delete shape1;
//     shape1 = NULL;
//     delete shape2;
//     shape2 = NULL;

//     system("pause");
//     return 0;
// }